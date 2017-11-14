package  main

import (
    "io"
    "fmt"
    "log"
    "net/http"
)

func sayHello(w http.ResponseWriter, r * http.Request) {
    fmt.Println("hello world!")
    io.WriteString(w,"<h1>hello world!</h1>")
    
}

func main() {
    http.HandleFunc("/",sayHello)
    er := http.ListenAndServe(":9090",nil)

    if er != nil{
        log.Fatal("ListenAndServe: ", er)
    }
}
