// 重複している行が含まれている全てのファイルを表示する
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//参照型の変数であるスライスやマップの場合は、ポインタ型を渡さなくてももとから参照渡しになっている

	counts := make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				panic("cannnot open file")
			}
			countLines(f, counts)
			f.Close()
		}
	}
	box := make([][]string, 0)
	for k, v := range counts {
		tmp := strings.Split(k, ":")
		tmp = append(tmp, strconv.Itoa(v))
		//fmt.Printf("key:%v file:%v v:%d\n", tmp[0], tmp[1], v)
		box = append(box, tmp)
	}
	for i, v := range box {
		for _, verb := range box[i+1:] {
			if v[0] == verb[0] {
				a, _ := strconv.Atoi(v[2])
				b, _ := strconv.Atoi(verb[2])
				fmt.Println("found duplicate:", v[0], v[1], verb[1], "count:", a+b, "(", a, "+", b, ")")
			}

		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	//バッファに全て貯めている状態
	input := bufio.NewScanner(f)
	//Scanで一列読み込む
	//input.Scan()
	//読み込んだものを表示する(バッファは読み込んだ分だけ消費される)
	//fmt.Println(input.Text())
	for input.Scan() {
		counts[input.Text()+":"+f.Name()]++
	}
}

//bufio.Reader()での読み込みが「指定した長さのバイト列ごと」なのに対して、これは「トークンごとの読み込み」をできるようにすることで利便性を向上させたものです。
// デフォルトでは改行をトークンとして文字を読み取っている
