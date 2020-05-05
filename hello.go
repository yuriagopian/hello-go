package main

	
import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func hello (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World !!!"))
}

func main() {
	r:= mux.NewRouter()
	r.HandleFunc("/", hello)
	http.ListenAndServe(":8081", r)
	fmt.Println("Hello world", r) 
   
}