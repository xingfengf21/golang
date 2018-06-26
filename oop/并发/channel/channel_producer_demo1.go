

func consumer(c chan int64) {
    var v int64
    ok := true

    for ok {
        if v, ok = <-c; ok {
            fmt.Println(v)
        }
    }
}


func producer(c chan int64, max int) {
    defer
    close(c)

    for i:= 0; i < max; i ++ {
        c <- time.Now().Unix()
    }
}

