// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package mail

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"testing"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func TestSend(t *testing.T) {

	testCases := []struct {
		a1 string
		a2 string
		a3 string

		e1 error
	}{
		{
			a1: "makarov-aleksei-z@yandex.ru",
			a2: "test case mbody",
			a3: "test case theme",

			e1: nil,
		},
	}

	for i, c := range testCases {
		if sErr := Send(c.a1, c.a2, c.a3); sErr != c.e1 {
			t.Fatalf(
				`
				Test failed:	TestSend (%d)
								want (%v), got (%v)
				`,
				i,
				c.e1,
				sErr,
			)
		}
	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
