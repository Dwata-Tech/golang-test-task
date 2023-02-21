package rabbitmq

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/Dwata-Tech/golang-test-task/database"
	"github.com/Dwata-Tech/golang-test-task/model"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"log"
)

var Channel *amqp.Channel

func PublishMessage(message []byte) error {
	logrus.Info("Producer: RabbitMQ")

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer func(connection *amqp.Connection) {
		err := connection.Close()
		if err != nil {
			logrus.Error("error: " + err.Error())
		}
	}(connection)

	logrus.Info("Successfully connected to RabbitMQ instance")

	// opening a channel over the connection established to interact with RabbitMQ
	channel, err := connection.Channel()
	if err != nil {
		logrus.Fatal("error: " + err.Error())
	}
	defer func(channel *amqp.Channel) {
		err := channel.Close()
		if err != nil {
			logrus.Error("error: " + err.Error())
		}
	}(channel)

	// declaring queue with its properties over the channel opened
	_, err = channel.QueueDeclare(
		"article", // name
		false,     // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // args
	)
	if err != nil {
		logrus.Fatal("error: " + err.Error())
	}
	//assigning into global variable
	Channel = channel
	//logrus.Info("Queue status:", queue)

	// publishing a message
	err = Channel.Publish(
		"",        // exchange
		"article", // key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)

	if err != nil {
		logrus.Error("error: " + err.Error())
		return err
	}
	logrus.Info("Successfully published message")
	return nil
}

func StartConsumer() {
	fmt.Println("Consumer: RabbitMQ")

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer func(connection *amqp.Connection) {
		err := connection.Close()
		if err != nil {

		}
	}(connection)

	fmt.Println("Successfully connected to RabbitMQ instance")

	// opening a channel over the connection established to interact with RabbitMQ
	channel, err := connection.Channel()
	if err != nil {
		logrus.Fatal("error: " + err.Error())
	}
	defer func(channel *amqp.Channel) {
		err := channel.Close()
		if err != nil {
			logrus.Error("error: " + err.Error())
		}
	}(channel)

	// declaring consumer with its properties over channel opened
	msgs, err := channel.Consume(
		"article", // queue
		"",        // consumer
		true,      // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       //args
	)
	if err != nil {
		logrus.Fatal("error: " + err.Error())
	}

	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {

			article := DecodeToArticle(msg.Body)
			res := database.Instance.Create(&article)
			if res.Error != nil {
				logrus.Error("Error:" + res.Error.Error())
			}
			logrus.Infof("Received Message: %v\n", article)
			logrus.Infof("Successfully saved to database, Article ID: %d", article.ID)
		}
	}()

	logrus.Info("Waiting for messages...")
	<-forever
}

func DecodeToArticle(s []byte) model.Article {

	p := model.Article{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal(err)
	}
	return p
}
