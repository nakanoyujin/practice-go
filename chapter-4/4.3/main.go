package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
	}
	reverse(slice)
	//pointerReverse((*[10]int)(slice))
	//スライス自体は参照渡し(中身の配列情報はポインタ)
	//rotateR(3, &slice)
	fmt.Println(slice)
}

// スライスはコンポジット型で配列へのポインタを渡している
// 配列へのポインタの中身を直接いじっているのでこれでsliceが書き換わる
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func pointerReverse(s *[10]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println("ポインタリバース")
	for i := 0; i < 10; i++ {
		fmt.Println(&s[i])
	}

}

// ref https://qiita.com/syoimin/items/9467ae884c53967e5b82
func rotateR(n int, slice *[]int) {
	sliceCopy := make([]int, len(*slice), len(*slice))
	//何個ぶんずれるか
	//2だったら8
	sliceNumber := len(*slice) - n
	for i := 0; i < n-1; i++ {
		sliceCopy[i] = (*slice)[sliceNumber+i+1]
	}
	for i := 0; i < sliceNumber+1; i++ {
		sliceCopy[i+n-1] = (*slice)[i]
	}
	*slice = sliceCopy
}
