package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		// emoi, err := strconv.Atoi(name)
		// if err != nil {
		// 	// handle error
		// }
		// r := rune(emoi)
		rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	})
	http.ListenAndServe(":8989", nil)
}
