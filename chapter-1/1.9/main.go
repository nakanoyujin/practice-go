package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, i := range os.Args[1:] {
		resp, err := http.Get(i)
		if err != nil {
			//Fが接頭の時は書き込み先を指定できる
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
		fmt.Println("StatusCode:", resp.Status)

	}
}
