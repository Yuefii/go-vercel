package handler

import "net/http"

func main() {
	http.HandleFunc("/api/hello", HelloWorld)
}
