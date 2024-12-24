package main

import("net/http")

func handler_err(w http.ResponseWriter, r *http.Request){
	respondwithErrors(w, 400, "Something went wrong")
	
}
