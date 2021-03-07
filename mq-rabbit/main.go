package main

import (
	server "github.com/henryinfanteg/heyadivinen_backgo/mq-rabbit/server"
)

func main() {
	var conection server.ConectionMQ

	conection.Username = "guest"
	conection.Password = "guest"
	conection.Host = "localhost"
	conection.Port = "5672"
	conection.QueueName = "leo"

	server.SendMsg(&conection, conection)
}
