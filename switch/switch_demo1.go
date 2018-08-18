package main

import "fmt"

func grade(score int) string  {
	g := ""
	switch  {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score:%d",score))
	case score < 60:
		g = "D"
	case score < 80:
		g = "C"
	case score <= 100:
		g = "A"

	}
    return g
}

func main() {

	fmt.Println(grade(100),grade(101))

}
