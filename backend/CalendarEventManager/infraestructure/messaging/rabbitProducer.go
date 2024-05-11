package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/cmd/config"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	headerCorrelationId   = "x-correlation-id"
	headerApplicationJson = "application/json"
	exchangeEvents        = "domainEvents"
)

type RabbitProducer struct {
	id   string
	lock *sync.Mutex
	conn *amqp.Connection
	ch   *amqp.Channel
	envs *config.Env
}

func NewProducer(envs *config.Env) *RabbitProducer {
	conn, err := amqp.Dial(fmt.Sprintf(`amqp://%s:%s@%s/%s`, envs.RABBITMQ_DEFAULT_USER, envs.RABBITMQ_DEFAULT_PASS, envs.RABBITMQ_HOST, envs.RABBITMQ_DEFAULT_VHOST))
	if err != nil {
		log.Panicln(err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	producerId := fmt.Sprintf(`%s-%s`, envs.APP_NAME_NOTIFICATIONS, uuid.New().String())
	log.Println("rabbit ready")
	return &RabbitProducer{envs: envs, conn: conn, lock: &sync.Mutex{}, id: producerId, ch: channel}
}

func (producer *RabbitProducer) PublishEvent(eventData Event, ctx context.Context) error {
	id, _ := uuid.NewRandom()
	var payload = Event{
		Name:    eventData.Name,
		EventId: eventData.EventId,
		Data:    eventData.Data,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = producer.ch.PublishWithContext(ctx,
		exchangeEvents, // exchange
		eventData.Name, // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			Headers: map[string]interface{}{
				"sourceApplication": producer.envs.APP_NAME_NOTIFICATIONS,
			},
			ContentType:     headerApplicationJson,
			ContentEncoding: "UTF-8",
			DeliveryMode:    2,
			Priority:        0,
			CorrelationId:   "",
			ReplyTo:         "",
			Expiration:      "",
			MessageId:       id.String(),
			Timestamp:       time.Now().UTC(),
			Type:            "",
			UserId:          "",
			AppId:           producer.envs.APP_NAME_NOTIFICATIONS,
			Body:            payloadBytes,
		})

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
