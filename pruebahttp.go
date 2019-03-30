package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

type Hello struct{}

func (h Hello) ServeHTTP(
		w http.ResponseWriter,
		r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	fichero, _ := os.OpenFile("fichero.txt", os.O_APPEND, 0777)
	nuevo_texto := os.Args[1]
	fichero.WriteString(nuevo_texto)
	fichero2, error := ioutil.ReadFile("fichero.txt")
	showError(error)
	fmt.Println(string(fichero2))
	var h Hello
	http.ListenAndServe("localhost:4000",h)	
}

func showError(e error){
	if (e != nil){
		panic(e)
	}
}