package main

import "fmt"

//函数作为值,类型

type testInt func(int) bool  //声明一个函数类型

func isOdd(integer int)bool  {
	if integer % 2 == 0{
		return false
	}
	return true
}

func isEvent(integer int)bool  {
	if integer % 2 == 0{
		return true
	}
	return false
}

//声明的函数类型在这个地方当做了一个参数
func filter(slice []int,f testInt) []int {
	var result []int
	for _,value := range slice{
		if f(value){
			result = append(result,value)
		}
	}
	return result
}


func main() {
	slice := []int {1,2,3,4,5,7}
	fmt.Println(slice)
	odd := filter(slice,isOdd)
	fmt.Println("Odd elements of slice are: ",odd)
	event := filter(slice,isEvent)
	fmt.Println("Even element of slice are: ",event)
}