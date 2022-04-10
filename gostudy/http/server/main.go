package main

import (
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		res.Write([]byte("我收到了！GET"))
	case "POST":
		//body, err := io.ReadAll(req.Body)
		header := res.Header()
		header["test"] = []string{"test1", "test2"}
		res.Write([]byte("我收到了！POST"))
	}

}

func main() {
	//http.HandleFunc("/test", handler)
	//http.ListenAndServe(":8080", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/test", handler)
	http.ListenAndServe(":8080", mux)
}
