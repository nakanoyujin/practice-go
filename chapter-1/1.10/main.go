package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	//ch = 0xc0001020c0みたいに構造体、スライス同様に参照渡し
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

var counter int
var mu sync.Mutex

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	//get url counter
	mu.Lock()
	count := counter
	counter++
	mu.Unlock()

	filename := fmt.Sprintf("%d%s", count, ".txt")

	//nbytes, err := io.Copy(io.Discard, resp.Body)
	file, err := os.Create("./tmp/" + filename)
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("While reading%s*%v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
