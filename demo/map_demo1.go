package main

import (
	"fmt"
)

func example(x int) int {
	if x == 0 {
		return 5
	} else { //必须在一行
		return x
	}
}

type PersionInfo struct {
	ID      string
	Name    string
	Address string
}

type MyMap interface {
	get()
}

func (persionInfo PersionInfo) get(persionInfo PersionInfo) string {

	return "666"

}

func get_personDB() map[string]PersionInfo {
	var personDB map[string]PersionInfo
	personDB = make(map[string]PersionInfo)

	personDB["1234"] = PersionInfo{"12345", "Tom", "Room 203"}
	personDB["1"] = PersionInfo{"1", "Jack", "Room 123"}

	return personDB
}

func main() {
	// var personDB map[string]PersionInfo
	// personDB = make(map[string]PersionInfo)

	// personDB["1234"] = PersionInfo{"12345", "Tom", "Room 203"}
	// personDB["1"] = PersionInfo{"1", "Jack", "Room 123"}

	// person, ok := personDB["1234"]
	personDB := get_personDB()
	fmt.Println(personDB)
	person, ok := personDB["1234"]

	// fmt.Println(reflect.TypeOf(
	//        personDB := get_personDB()
	//        ))
	va := personDB["1234"]
	fmt.Println("va", va)
	fmt.Println("ok", ok)

	if ok {
		fmt.Println("Found Person", person.Name, person.Address)
	} else {
		fmt.Println("Did not find person with ID 1234")
	}
	fmt.Println(example(2))
}
