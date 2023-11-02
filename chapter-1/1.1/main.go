// echo
//  プログラムを起動したコマンド名を表示する
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
	s := ""
	sep := " "
	fmt.Printf("%v:", os.Args[0])
	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	fmt.Println(s)
}
