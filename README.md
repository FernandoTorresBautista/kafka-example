# Kafka 
Test to setup kafka and produce/consume messages.

## docker-composer
You just need to do the docker compose up to be able to use kafka.

## api
The api has the entry for produce messages using the /producer/producer.go <br> 
The consumer is initialzed after the api to consume the messages produced /consumer/consumer.go <br>

## .json
kafka-test.postman_collection.json has a postman service to produce a message

it's necessary to have installed docker and go
