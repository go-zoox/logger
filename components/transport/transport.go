package transport

import "github.com/go-zoox/logger/components/message"

type Transport interface {
	Write(message *message.Message)
}
