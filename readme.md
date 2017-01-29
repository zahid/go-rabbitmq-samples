## This is my path through the [RabbitMQ tutorials](https://www.rabbitmq.com/getstarted.html)

#### Pre-requisites:

This contains a `docker-compose.yml` that runs a RabbitMQ server.

```
docker-compose up -d
```

All the golang files can be executed via `go run...`.

#### Tutorial 1: Hello World

![](https://www.rabbitmq.com/img/tutorials/python-one-overall.png)
> https://www.rabbitmq.com/tutorials/tutorial-one-python.html

Send a message to the queue:
```
$ go run tutorial-1/send.go
```

You can check out what's happening with the queues between `send.go` and `receive.go` programs by listing the queues:

```
$ docker-compose exec rabbitmq rabbitmqctl list_queues
Listing queues ...
hello	1
```

Then receive from the queue:
```
$ go run tutorial-1/receive.go
2017/01/28 23:07:45  [*] Waiting for messages. To exit press CTRL+C
2017/01/28 23:07:45 received a message : Hello there!
```



