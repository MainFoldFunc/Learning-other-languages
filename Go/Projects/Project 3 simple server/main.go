package main

import(
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not foumd", http.StatusNotFound)
		return 
	}
	if r.Method != "GET"{
		http.Error(w, "Method is not suported", http.StatusNotFound)
		return
	}
	fmt.Println(w, "Hello world")
}

func formHandler (w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		fmt.Printf(w, "ParseForm() error: %v"err)
	}
	fmt.Fprintf(w, "POST request succesful")
	name := r.FormValue("name")
	adress := r.FormValue("adress")
	fmt.Fprintf(w, "Name: %v\nAdress: %v", name, adress)
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Strting the server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal(err)
	}



}

