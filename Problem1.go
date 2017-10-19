package main

import (
    "io"
    "net/http"
    "log"
    "fmt"
)

func DrawMenu(w http.ResponseWriter){
	Header().Set("Content-Type", "text/html")
	io.WriteString("<h1> Guessing Game <h1/>" + "\n")
    
}
func main() {
	
	DrawMenu( )
    fmt.Println("listen on 3000")
    err := http.ListenAndServe(":3000", nil)

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}