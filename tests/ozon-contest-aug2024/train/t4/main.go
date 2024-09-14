package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func truncate(obj interface{}) interface{} {
	arr, ok := obj.([]any)
	if ok {
		var res []interface{}
		for _, v := range arr {
			v = truncate(v)
			if !(len(arr) == 0) {
				res = append(res, v)
			}
		}
		return res
	}

	dct, ok := obj.(map[string]any)
	if ok {
		for k, v := range dct {
			dct[k] = truncate(v)
			if dct[k] == nil {
				delete(dct, k)
			}
		}
		return dct
	}

	return obj
}

func prettify(s []string) interface{} {
	var (
		i interface{}
		d = strings.Join(s, "")
	)

	if err := json.Unmarshal(
		[]byte(d), &i,
	); err != nil {
		panic(err)
	}
	return truncate(i)
}

func main() {
	ri := bufio.NewReader(os.Stdin)

	wo := bufio.NewWriter(os.Stdout)
	defer wo.Flush()

	var r []any

	var t int
	fmt.Fscan(ri, &t)

	for i := 0; i < t; i++ {
		var n int
		var f []string
		fmt.Fscan(ri, &n)

		for j := 0; j <= n; j++ {
			// var s string
			s, err := ri.ReadString('\n')
			if err != nil {
				panic(err)
			}
			f = append(
				f,
				strings.Join(strings.Fields(s), " "),
			)
		}

		r = append(r, prettify(f))
	}

	m, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(m))
}
