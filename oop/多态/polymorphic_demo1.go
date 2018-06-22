package main 


import(
	"fmt"
)

type user struct{
	name string 
	email string  
}

type admin struct{
	person user 
	level int 
}

func (u *user) notify(){
	fmt.Printf("Sending user email to %s<%s>\n",u.name,u.email)
}

func (a *admin) notify(){
	fmt.Printf("Sending user email to %s<%s>\n",a.person.name,a.person.email)
}

type notifier interface{
	notify()
}


func sendNotification(n notifier){
	n.notify()
}

func main (){
	bill := user{"Bill","bill@email.com"}
	sendNotification(&bill)
	admin := admin{
		person:user{
			name:"admin",
			email:"admin@email.com",
		},
		level:12,
	}
	sendNotification(&admin)
}