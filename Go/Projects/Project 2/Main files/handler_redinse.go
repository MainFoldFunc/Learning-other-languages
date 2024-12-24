package main

import("net/http")

func handlerReadiness(w http.ResponseWriter, r *http.Request){
	respondpithJSON(w, 200, struct{}{})
	
}
