package main

import (
	"go-2025/handler"
)

func main() {
	// fmt.Println(handler.ToCamelCase("the-stealth-warrior"))
	// fmt.Println(handler.ToCamelCase("The_Stealth_Warrior"))
	// fmt.Println(handler.ToCamelCase(""))

	// arr := [...][]int{
	// 	{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5, 5},
	// 	{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5},
	// 	{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5},
	// 	{10},
	// 	{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1},
	// 	{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10},
	// }
	// // sol := [...]int{5, -1, 5, 10, 10, 1}
	// for _, v := range arr {
	// 	cap_v := v
	// 	fmt.Println("oo", handler.FindOdd(cap_v))
	// }
	// f := handler.Fibonacci()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }

	// fmt.Println(handler.Fibonacci2(10))
	handler.Point()

	// r := gin.Default()

	// r.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, "pong")
	// })

	// r.Run()
}
