package sender

import (
	"context"
	"testing"
)

func TestSender_SendEmail(t *testing.T) {
	var e = 0

	// немного отме
	for i := 0; i < 100; i++ {
		s := New()
		if err := s.SendEmail(context.TODO(), "", ""); err != nil {
			e++
		}
	}

	if e == 0 {
		t.Fatalf(
			"impossible",
		)
	}
}
