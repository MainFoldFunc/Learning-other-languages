package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func respondpithJSON(p http.ResponseWriter, code int, payload interface{}){
	data, err := json.Marshal(payload)
	if err != nil{
		log.Println("Failed to marshal JSON response: \n%v", payload)
		p.WriteHeader(500)
		return 

	}
	p.Header().Add("Content-type", "application/json")
	p.WriteHeader(code)
	p.Write(data)
}


func respondwithErrors(p http.ResponseWriter, code int, msg string){
	if code > 499{
		log.Println("Responding with 500 level error: ", msg)
	}
	type errorResponse struct{
		Error string "json: error"
	}

	respondpithJSON(p, code, errorResponse{
		Error: msg,
	})
	
}

