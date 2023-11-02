package main

import (
	"fmt"
	"strings"
	"testing"
)

// func main() {
// 	start := time.Now()
// 	fmt.Println(strings.Join(os.Args[1:], " "))
// 	passed := time.Since(start).Seconds()
// 	fmt.Println("time passed:", passed)

// 	start = time.Now()
// 	var s, sep string
// 	for i := 1; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// 	passed = time.Since(start).Seconds()
// 	fmt.Println("time passed:", passed)
// }

var args = []string{"cmd", "geg", "ae", "gaeewa"}

func BenchmarkA(b *testing.B) {
	fmt.Println(strings.Join(args[1:], " "))
}

func BenchmarkB(b *testing.B) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

// -benchmem
// メモリ関連の情報を出力。

// -benchtime t
// 計測に使う時間をtで指定する。defaultは1s

// -cpuprofile=*.prof
// 詳細なCPUプロファイルが取れる go tool pprofで内容を見る

// -count
// テストする回数を指定

// -cpu
// 実行するcpu数

// geg ae gaeewa
// goos: darwin
// goarch: amd64
// pkg: 1.2
// cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
// BenchmarkA-4    geg ae gaeewa
// geg ae gaeewa
// geg ae gaeewa
// geg ae gaeewa
// geg ae gaeewa
// 1000000000               0.0001239 ns/op
// geg ae gaeewa
// BenchmarkB-4    geg ae gaeewa
// geg ae gaeewa
// geg ae gaeewa
// geg ae gaeewa
// geg ae gaeewa
// 1000000000               0.0000128 ns/op
// PASS
// ok      1.2     0.321s

// 実行した回数
// １回あたりの実行に掛かった時間(ns/op)
// １回あたりのアロケーションで確保した容量(B/op)
// 1回あたりのアロケーション回数(allocs/op)
