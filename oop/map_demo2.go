
package main

import (
	"fmt"
	std "fmt"
)

type CcMap map[string] string 

//为Map增加Get方法
func (c CcMap) Get(key string) string{
	value,ok := c[key]
	if ok{
		return value
	}else{
		return ""
	}
}

//为Map增加Keys方法,获取所有的key
func (c CcMap) Keys() [] string {
	var s [] string 
	for k := range c{
		s = append(s,k)
	}
	return s 
}

//为Map增加Values方法,获取所有的值
func (c CcMap) Values() [] string {
	var s [] string 
	for k := range c{
		s = append(s,c[k])
	}
	return s 
}

func main(){
	// var cityMap map[string] string 
	// cityMap = make(map[string] string)
	var cityMap CcMap
	cityMap = make(CcMap)
	cityMap["x"] = "西安"
	cityMap["b"] = "北京"
	cityMap["c"] = "重庆"
	std.Println(cityMap)

	//获取value 
	val,ok :=cityMap["x2"]
	if ok{
		fmt.Println(val)
	}else{
		fmt.Println("Not Found.")
	}
	value := cityMap.Get("b")
	fmt.Println(value)
	fmt.Println(cityMap.Get("no"))

	//遍历map
	for city := range cityMap{
		fmt.Println("city",city,cityMap[city])
	}
	fmt.Println(cityMap.Keys())
	fmt.Println(cityMap.Values())
}