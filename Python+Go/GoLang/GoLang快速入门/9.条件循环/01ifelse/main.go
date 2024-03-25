package main

func main() {
	//// 1、if else
	//score := 65
	//if score > 90 {  // 条件不成立（走else语句）
	//	fmt.Println("高分")
	//} else {
	//	fmt.Println("小于90") // 小于90
	//}

	////// 2、score写在if后面，这个score变量就只能在if代码块使用
	//if score := 65; score > 90 {
	//	fmt.Println("高分")
	//} else {
	//	fmt.Println(score)
	//	fmt.Println("小于90") // 小于90
	//}
	////fmt.Println(score)  // score在if语句内定义，不能在外部引用

	//// 3、if else if .... else
	//// if语句特点：从上向下匹配，成立就返回，而不会匹配后面的条件
	//score := 66
	//if score > 90 {
	//	fmt.Println("优秀")
	//} else if score > 80 {
	//	fmt.Println("良好")
	//} else if score > 70 {
	//	fmt.Println("一般良好")
	//} else {
	//	fmt.Println("一般")
	//}

}
