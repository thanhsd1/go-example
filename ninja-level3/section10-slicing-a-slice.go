package main

import "fmt"

func main() {
	x := []int{5, 4, 3, 1, 8}
	fmt.Println(x[1])  // lấy ra 1 phần tử trong mảng
	fmt.Println(x)     // in ra toàn bộ mảng của biên x
	fmt.Println(x[2:]) // xoá đi  phần tử trong mảng. Hàm hiện tại là xoá đi 2 phần tử đếm từ trái sang phải. VD: "[2:]" ngược lại thì "[:2]" xoá từ phải sang trái
}
