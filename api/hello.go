package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

// func main() {
// 	http.HandleFunc("/", greet)
// 	http.ListenAndServe(":8080", nil)
// }
