package user_service

import (
	"context"
	"encoding/json"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateRabbit() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@queue:5672/")
	if err != nil {
		return nil
	}
	return conn
}

func Send(w http.ResponseWriter, r *http.Request, statuscode *int) {
	conn := CreateRabbit()
	if conn == nil {
		return
	}
	ch, err := conn.Channel()
	if err != nil {
		return
	}
	q, err := ch.QueueDeclare(
		"logger",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}
	ch.PublishWithContext(context.Background(), "", q.Name, false, false, amqp.Publishing{
		ContentType: "test/plain",
		Body:        Encode(w, r, *statuscode),
	})
}
func SendRabbit(w http.ResponseWriter, r *http.Request, statuscode *int) {
	go func() {
		Send(w, r, statuscode)
	}()
}

func Encode(w http.ResponseWriter, r *http.Request, statuscode int) []byte {
	data := make(map[string]interface{})
	data["Method"] = r.Method
	data["Url"] = r.URL.String()
	data["Host"] = r.Host
	data["Status"] = statuscode
	bin, _ := json.Marshal(data)
	return bin
}
