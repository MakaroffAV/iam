package db

import "testing"

func TestConn(t *testing.T) {
	_, err := Conn()
	if err != nil {
		t.Fatalf(
			"got: %v; want: %v; \n", err, nil,
		)
	}
}
