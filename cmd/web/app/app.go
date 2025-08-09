package application

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Port string
}

const defaultPort = ":4000"

func New() *Application {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	} else if port[0] != ':' {
		port = ":" + port
	}
	return &Application{Port: port}
}

func (app *Application) StartServer() {
	// router
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	//start the server
	go func() {
		log.Printf("Starting server at port %s", app.Port)
		err := http.ListenAndServe(app.Port, mux)
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "Hello World")
}
