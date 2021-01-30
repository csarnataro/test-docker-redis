package messages

// Message is a message
type Message struct {
	Endpoint Endpoint            `json:"endpoint"`
	Data     []map[string]string `json:"data"`
}
