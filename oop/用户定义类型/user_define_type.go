package main 

import (
	"fmt"
)

type user struct{
	name string 
	emil string 
	ext int 
	privileged bool 
}

type admin struct{
	person user 
	level string 
}

func main(){
	lisa := user{
		name:"Lisa",
		emil:"Lisa@email.com",
		ext:123,
		privileged:false,
	}
	fmt.Println(lisa)
	
	fred := admin{
		person:user{
			name:"admin",
			emil:"admin@emil.com",
			ext:12,
			privileged:true,
		},
		level:"super",
	}
	fmt.Println(fred)
	fmt.Println(fred.person.emil)
}