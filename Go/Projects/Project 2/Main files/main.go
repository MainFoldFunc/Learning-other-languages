package main

import ("fmt"
	"os"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors")
func main(){
	//var exit string
	godotenv.Load()
	
	portStr := os.Getenv("PORT")
	if portStr == ""{
		log.Fatal("PORT not found...")
	}
	fmt.Println("Port: ", portStr)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:	[]string{"https://*", "http://*"},
		AllowedMethods:	[]string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:	[]string{"*"},
		ExposedHeaders:	[]string{"Link"},
		AllowCredentials: false,
		MaxAge:	300,
	}))

	routerv1 := chi.NewRouter()
	routerv1.Get("/ready", handlerReadiness)
	routerv1.Get("/err", handler_err)

	router.Mount("/v1", routerv1)

	server := &http.Server{
		Handler: router,
		Addr: ":" + portStr,
	}

	log.Printf("Server starting at port: %v", portStr)

	err := server.ListenAndServe()

	if err != nil{
		log.Fatal(err)
	}


}
