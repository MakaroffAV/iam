package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func walk(p string, wg *sync.WaitGroup, fileSizes chan<- int64, fileErrors chan<- string) {
	defer wg.Done()

	l, lErr := os.ReadDir(p)
	if lErr != nil {
		fileErrors <- lErr.Error()
	}

	for _, e := range l {
		if e.IsDir() {
			wg.Add(1)
			go walk(filepath.Join(p, e.Name()), wg, fileSizes, fileErrors)
		} else {
			s, sErr := os.Stat(filepath.Join(p, e.Name()))
			if sErr != nil {
				fileErrors <- sErr.Error()
				continue
			}
			fileSizes <- s.Size()

		}
	}

}

func main() {

	var sum int64

	roots := []string{"/Users/makarov_aleksei/Desktop"}
	var wg sync.WaitGroup
	fileSizes := make(chan int64)
	fileErrors := make(chan string)
	for _, root := range roots {
		wg.Add(1)
		go walk(root, &wg, fileSizes, fileErrors)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
		close(fileErrors)
	}()

	for {
		select {
		case s, ok := <-fileSizes:
			{
				if ok {
					sum += s
				} else {
					fileSizes = nil
				}
			}
		case e, ok := <-fileErrors:
			{
				if ok {
					fmt.Fprintf(os.Stderr, "error:%v\n", e)
				} else {
					fileErrors = nil
				}
			}
		}
		if fileSizes == nil && fileErrors == nil {
			break
		}
	}

	fmt.Println(float64(sum) / float64(1e+9))

}
