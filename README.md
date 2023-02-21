# Golang Test Task

The project is about simple 5 APIs as below. 
1. Post /api/article
2. GET /api/articles
3. GET /api/article/{id}
4. POST /api/comment
5. GET /api/comment/{article_id}

It uses the RabbitMQ for posting article to database. Where producer receives the request from REST API and when consumer receives the messages it saves it into database.


# Flow Diagram 

[![](https://mermaid.ink/img/pako:eNp1kk1rwzAMhv-KMIRskF0GveRQaJsedgjkY7BLLo6jpobE6WyZbYT-99lNQlfaKWBiPa-EZb8jE0ODLGZBMEolKYYRQjpij2EcdrI9UhhBaPDTohIuB2NoDab8-0M2dAzjA-8MnsF9QVCpRZdI3mreVwpccEuDsn2Net4LGjRsgRvY6uHLLPkT1ySFPHFFUHha7Mt32GRv9zjzONNDY8Wj6jS_lPO6lpTm93zn8W5QxvaPyhOPE0685gYrNQk0CgLd1k-vq1UEy_I8QR_bl_W6iCEbDMHGN-vwCgsHMwdt3UlzhBSN4e0fnjme5ldBPwmAhnkMWObwkeZTuwLJagWlFcLpb7sV_2N_mO3_-NJ9N0-SXg9ye2E-dk6YxFC698RlZq-83p1XoWpYxFxhz2XjnDb6dMUuJqtY7H4bPHDbUcUqdXZSe2o44b6Rri-LSVuMmDdR-aPEsp80s8_YbMNfbf3iGA?type=png)](https://mermaid.live/edit#pako:eNp1kk1rwzAMhv-KMIRskF0GveRQaJsedgjkY7BLLo6jpobE6WyZbYT-99lNQlfaKWBiPa-EZb8jE0ODLGZBMEolKYYRQjpij2EcdrI9UhhBaPDTohIuB2NoDab8-0M2dAzjA-8MnsF9QVCpRZdI3mreVwpccEuDsn2Net4LGjRsgRvY6uHLLPkT1ySFPHFFUHha7Mt32GRv9zjzONNDY8Wj6jS_lPO6lpTm93zn8W5QxvaPyhOPE0685gYrNQk0CgLd1k-vq1UEy_I8QR_bl_W6iCEbDMHGN-vwCgsHMwdt3UlzhBSN4e0fnjme5ldBPwmAhnkMWObwkeZTuwLJagWlFcLpb7sV_2N_mO3_-NJ9N0-SXg9ye2E-dk6YxFC698RlZq-83p1XoWpYxFxhz2XjnDb6dMUuJqtY7H4bPHDbUcUqdXZSe2o44b6Rri-LSVuMmDdR-aPEsp80s8_YbMNfbf3iGA)