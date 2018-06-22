package main 
import "fmt"

type Header map[string] [] string 

func (h Header) Get(key string) string{
	value, ok := h[key]
	if ok{
		return "2"

	}else{
		return ""
	}
}

func main(){

}