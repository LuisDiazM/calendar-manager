package messaging

type Event struct {
	Name    string      `json:"name"`
	EventId string      `json:"eventId"`
	Data    interface{} `json:"data"`
}
