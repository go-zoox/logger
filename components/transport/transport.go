package transport

type Transport interface {
	// Write(message *message.Message)
	Write(p []byte) (n int, err error)
	WriteWithLevel(p []byte, level string) (n int, err error)
}
