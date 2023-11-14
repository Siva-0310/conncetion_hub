package logger

import (
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Log struct {
	Method string `json:"method"`
	Url    string `json:"url"`
	Status int    `json:"status"`
	Host   string `json:"Host"`
}

func RabbitConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@queue:5672/")
	if err != nil {
		return nil
	}
	return conn
}
func Main() {
	conn := RabbitConnection()
	if conn == nil {
		return
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return
	}
	defer ch.Close()
	q, err := ch.QueueDeclare("logger", false, false, false, false, nil)
	if err != nil {
		return
	}
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return
	}
	for i := range msgs {
		log := &Log{}
		json.Unmarshal(i.Body, log)
		fmt.Println(log)
	}
}
