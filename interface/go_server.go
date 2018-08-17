package main
import (
	"fmt"
	"log"
	"net/http"
	"strings"
)
func sayhelloName(w http.ResponseWriter,r *http.Request){ //访问主体回应
	r.ParseForm() //解析url传递的参数,对于POST则解析响应包的主体
	//如果没有调用ParseForm方法则无法获取表单的数据.
	fmt.Println(r.Form)//输出url的参数,到服务器上
	fmt.Println(r.URL.Path)//输出url的路径
	fmt.Println(r.URL.Scheme)//输出方案
	for k,v :=range r.Form{
		fmt.Println(k)
		fmt.Println(strings.Join(v,""))//加入字符串输出看效果
	}
	fmt.Fprintf(w,"hello world")//回应内容
}
func login(w http.ResponseWriter,r *http.Request){
	fmt.Println(r.Method)//输出客户端请求的方法
	if r.Method=="GET"{
		//t,_:=template.ParseFiles("login.html") //解析html代码,就是那个表单
		//log.Println(t.Execute(w,nil))
		log.Println(22222222)
	}else{
		r.ParseForm()
		fmt.Println(r.Form["username"])  //字典获取username的值
		fmt.Println(r.Form["password"])
	}
}

func main(){
	log.Println("main....")
	http.HandleFunc("/",sayhelloName)  //设置访问的路由
	http.HandleFunc("/login",login)
	err:=http.ListenAndServe(":9000",nil)//设置监听端口
	if err !=nil{
		log.Fatal("listenandserver",err)//报错则输出错误并退出
	}
}
func B2S(bs []uint8) string {
	ba := make([]byte,0, len(bs))
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}
