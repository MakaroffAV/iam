// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package app

import (
	"fmt"
	"parser_ecom/pkg/req"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func categories(storeId string) {

	var s = []string{
		"https://sbermarket.ru/api/v3/stores/8391/categories",
	}

	for len(s) > 0 {

		p := s[0]
		s = s[1:]

		r, rErr := req.Get(p)
		if rErr != nil {
			fmt.Println(rErr)
		}

		fmt.Println(string(r))

	}

}

func Run(storeURL string) {

}
