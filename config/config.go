package config

import "strconv"

var (
	// Server is the MQTT server address
	Server string

	// Port is the MQTT server listen port
	Port int

	// ClientID is how the name of the client
	ClientID string

	// TopicPrefix is just that, a prefix for the presented data keys
	TopicPrefix string

	// Interval is the poll interval in seconds
	Interval int
)

// ConnectionString returns the MQTT connection string
func ConnectionString() string {
	return "tcp://" + Server + ":" + strconv.Itoa(Port)
}
