package main 


import(
	"fmt"
)

type user struct{
	name string 
	email string  
}

type admin struct{
    user 
	level int 
}

func (u *user) notify(){
	fmt.Printf("Sending user email to %s<%s>\n",u.name,u.email)
}

func (a *admin) notify(){
	fmt.Printf("Sending user email to %s<%s>\n",a.user.name,a.user.email)
}

type notifier interface{
	notify()
}


func sendNotification(n notifier){
	n.notify()
}

func main (){

	admin := admin{
		user:user{
			name:"admin",
			email:"admin@email.com",
		},
		level:12,
	}
	admin.notify()
	admin.user.notify()
}