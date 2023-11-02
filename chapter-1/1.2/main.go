package main

import (
	"fmt"
	"os"
)

// func main() {
// 	var s, sep string
// 	//OSの引数は1から始まる
// 	//0はプロセスID
// 	for i := 1; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Printf("%d : %v\n", index, arg)
	}
}
