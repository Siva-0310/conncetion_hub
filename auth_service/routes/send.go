package routes

import (
	"context"
	"encoding/json"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

func RabbitConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@queue:5672/")
	if err != nil {
		return nil
	}
	return conn
}
func RabbitChannel() {

}
func Send(w http.ResponseWriter, r *http.Request) {
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
	ch.PublishWithContext(context.Background(), "", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        Encode(w, r),
	})
}

func Encode(w http.ResponseWriter, r *http.Request) []byte {
	data := make(map[string]interface{})
	data["Method"] = r.Method
	data["Url"] = r.URL.String()
	if sc, ok := w.(interface{ StatusCode() int }); ok {
		data["ResponseStatus"] = sc.StatusCode()
	} else {
		data["ResponseStatus"] = http.StatusOK
	}
	bin, _ := json.Marshal(data)
	return bin
}
