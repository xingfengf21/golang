package main 
import  std "fmt"

func Range(start int ,end int ,step int) [] int {
	var list [] int 
	if start > end{
		return list 
	}
	for t := start; t < end; t+=step {
		if t < end{
			list = append(list,t )
		} 
	}

	return list 
}

func main(){
	std.Println(Range(0,10,2))
}
