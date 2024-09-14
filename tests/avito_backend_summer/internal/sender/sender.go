package sender

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

type Sender struct{}

func New() *Sender {
	return &Sender{}
}

func (s *Sender) SendEmail(ctx context.Context, recipient string, message string) error {
	// Имитация отправки сообщения
	// немного изменил время ожидания, чтобы не было таймаута теста
	duration := time.Duration(rand.Int63n(10)) * time.Millisecond
	time.Sleep(duration)

	// Имитация неуспешной отправки сообщения
	errorProbability := 0.1
	if rand.Float64() < errorProbability {
		return errors.New("internal error")
	}

	// зачем????
	// fmt.Printf("send message '%s' to '%s'\n", message, recipient)

	return nil
}
