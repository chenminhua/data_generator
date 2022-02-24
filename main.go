package main

import "github.com/chenminhua/data_generator/es"

func main() {
	// 生成1000000万用户
	//ch := make(chan int)
	//constant.LoadNamesData()
	//mysql.OpenDB()
	//for i := 0; i < 32; i++ {
	//	go func() {
	//		for {
	//			idx := <-ch
	//			println(idx)
	//			mysql.CreateUser(idx, idx+100)
	//		}
	//	}()
	//}
	//for i := 0; i <= 1000000; i += 100 {
	//	println(i)
	//	ch <- i
	//}
	//close(ch)
	es.Run()
}
