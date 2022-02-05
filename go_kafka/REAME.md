# Go_Kafka

## Technologies Used

- Docker version 20.10.12, build e91ed57
- Docker Compose version v2.2.3
- Go version go1.17.5 linux/amd64
- [Sarama]("https://github.com/Shopify/sarama") v1.31.1

## Goals
- [X] Consume topic message
- [ ] Post message in topic

## How to Test
1. Up Zookeeper and Kafka
```
docker-compose up -d
```

2. Create topic `topic-test`

a\) By command line
```
kafka --create --topic topic-test --replication-factor 1 --partitions 3
```

b\) By auxiliary software

I suggest using [Conduktor](https://www.conduktor.io/download)

3. Enable one of the two functions `postMessageInTopic` or `postMessageInTopic`

`postMessageInTopic` to get topic messages

`postMessageInTopic` to post to messages in the topic