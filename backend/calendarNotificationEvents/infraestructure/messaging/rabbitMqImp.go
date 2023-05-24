package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/cmd/config"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"golang.org/x/sync/errgroup"
)

type EventHandlerFunc func(context.Context, map[string]interface{}) error

type messageData struct {
	Resource string      `json:"resource"`
	Name     string      `json:"name"`
	Data     interface{} `json:"data"`
}

type RabbitMQConsumer struct {
	conn          *amqp091.Connection
	env           *config.Env
	eventHandlers map[string]EventHandlerFunc
}

func NewRabbitMQConsumer(environments *config.Env) *RabbitMQConsumer {
	conn, err := amqp091.Dial(fmt.Sprintf(`amqp://%s:%s@%s/%s`, environments.RABBITMQ_DEFAULT_USER,
		environments.RABBITMQ_DEFAULT_PASS,
		environments.RABBITMQ_HOST,
		environments.RABBITMQ_DEFAULT_VHOST))
	if err != nil {
		log.Panicln(err)
	}
	return &RabbitMQConsumer{env: environments, conn: conn, eventHandlers: map[string]EventHandlerFunc{}}
}

func (amqpRabbit *RabbitMQConsumer) Start() error {
	channel, err := amqpRabbit.conn.Channel()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	err = channel.Qos(8, 0, false)
	if err != nil {
		log.Println(err)
		return err
	}
	queue, err := amqpRabbit.CreateQueue(channel, amqpRabbit.eventHandlers)
	if err != nil {
		log.Println(err)
		return err
	}
	tag := fmt.Sprintf("%s-%s", amqpRabbit.env.APP_NAME, uuid.NewString())
	msgs, err := channel.Consume(
		queue.Name, // queue
		tag,
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Println(err)
		return err
	}
	semaphore := make(chan bool, 4)

	g, gCtx := errgroup.WithContext(context.TODO())

	for msg := range msgs {
		delivery := msg
		semaphore <- true
		g.Go(func() error {
			return amqpRabbit.handleMessage(gCtx, delivery, semaphore)
		})
	}
	<-gCtx.Done()
	if !channel.IsClosed() {
		err = channel.Cancel(tag, false)
		if err != nil {
			log.Println(err)
			return err
		}
		err = channel.Close()
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return g.Wait()
}

func (amqpRabbit *RabbitMQConsumer) AddEventHandler(event string, handler EventHandlerFunc) {
	amqpRabbit.eventHandlers[event] = handler
}

func (amqpRabbit *RabbitMQConsumer) CreateQueue(channel *amqp091.Channel, eventHandlers map[string]EventHandlerFunc) (*amqp091.Queue, error) {
	err := channel.ExchangeDeclare(
		eventExchange, // name
		"topic",       // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return nil, err
	}
	eventsQueue, err := channel.QueueDeclare(
		fmt.Sprintf(`%s.%s`, amqpRabbit.env.APP_NAME, eventSuffix), //name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}
	for event := range eventHandlers {
		err = channel.QueueBind(eventsQueue.Name, event, eventExchange, false, nil)
		if err != nil {
			return nil, err
		}
	}
	return &eventsQueue, nil
}

func (c *RabbitMQConsumer) handleMessage(cx context.Context, delivery amqp091.Delivery, semaphore chan bool) error {
	defer func() {
		<-semaphore
	}()

	var mdata messageData

	err := json.Unmarshal(delivery.Body, &mdata)
	if err != nil {
		log.Println(err)
		err = delivery.Nack(false, true)

		if err != nil {
			log.Println(err)
		}

		return err
	}
	handler, exists := c.eventHandlers[mdata.Name]
	if !exists {
		log.Println(err)
		err = delivery.Nack(false, true)

		if err != nil {
			log.Println(err)
		}

		return err
	}

	body, _ := mdata.Data.(map[string]interface{})

	err = handler(cx, body)
	if err != nil {
		log.Println(err)
		err = delivery.Nack(false, true)

		if err != nil {
			log.Println(err)
		}

		return err
	}

	err = delivery.Ack(false)

	if err != nil {
		log.Println(err)

		return err
	}

	return nil
}
