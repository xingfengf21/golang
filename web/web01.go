package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter,r * http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Scheme)
	fmt.Println(r.Form["url_log"])
	for k,v := range r.Form{
		fmt.Println("key",k)
		fmt.Println("val",strings.Join(v,""))
	}
	fmt.Println(w,"Hello astaxie")
    log.Println("66666")
}

func main() {
	http.HandleFunc("/",sayhelloName)
	err := http.ListenAndServe(":9000",nil)
	if err != nil{
		log.Fatal("ListenAndServer:",err)
	}
}