package main 
import (
	"fmt"
	"runtime"
	"sync"
)
var (
	counter int
	wg sync.WaitGroup 
	//用mutex定义一段代码临界区

	mutex sync.Mutex
)

func main(){
	wg.Add(2)
	
	go incCounter(1)
	go incCounter(2)

	wg.Wait()
    fmt.Println("Final Counter:",counter)
}

func incCounter(id int){
	defer wg.Done()
	
	for count:=0;count<2;count++{
		mutex.Lock()
		{
			value :=counter
			runtime.Gosched()

			value++
			counter = value
		}
		mutex.Unlock()
	}
}

