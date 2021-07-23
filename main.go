package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"math/rand"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/rs/cors"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {
	//opzione1()
	//opzione2()
	opzione3()
}

func opzione1() {
	var port int
	flag.IntVar(&port, "port", 8080, "La porta in ascolto è")
	flag.Parse()

	// Aggiungiamo CORS middleware per accedere al serever Vue
	// attivo su http://localhost:8080
	corsMiddleware := cors.New(cors.Options{AllowedOrigins: []string{"http://localhost:8080"}})
	http.Handle("/api/v1/embed", corsMiddleware.Handler(http.HandlerFunc(getFrasiAPI)))

	stripped, err := fs.Sub(frontend, "frontend/dist")
	if err != nil {
		log.Fatalln(err)
	}

	http.Handle("/", http.FileServer(http.FS(stripped)))
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func opzione2() {
	var port int
	flag.IntVar(&port, "port", 8080, "La porta in ascolto è")
	flag.Parse()

	stripped, err := fs.Sub(frontend, "frontend/dist")
	if err != nil {
		log.Fatalln(err)
	}

	frontendFS := http.FileServer(http.FS(stripped))
	http.Handle("/", frontendFS)

	http.Handle("/api/v1/embed", http.HandlerFunc(getFrasiAPI))

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// opzione3 permette al server Go di servire direttamente gli asset del frontend.
// Questo richiede che gli asset del fontend siano compilati ogni volta ci sia una modifica
// attraverso il comando `yarn watch`.
func opzione3() {
	http.Handle("/api/v1/embed", http.HandlerFunc(getFrasiAPI))

	// L'implementazione di questa funzione sarà gestita dal codice presente nei
	// files fs_dev.go e fs_prod.go, in base al build tag `prod` che verrà usato.
	frontend := getFrontend()

	http.Handle("/", http.FileServer(http.FS(frontend)))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

type Law struct {
	Titolo    string `json:"titolo,omitempty"`
	Enunciato string `json:"enunciato,omitempty"`
}

var FrasiAPI = []Law{
	{
		Titolo:    "Frase 1",
		Enunciato: "Questa è la frase 1, un esempio di risposta.",
	},
	{
		Titolo:    "Frase 2",
		Enunciato: "Questa è la frase 2, un esempio di risposta.",
	},
	{
		Titolo:    "Frase 3",
		Enunciato: "Questa è la frase 3, un esempio di risposta.",
	},
}

func getFrasiAPI(w http.ResponseWriter, r *http.Request) {
	randomFrasi := FrasiAPI[rand.Intn(len(FrasiAPI))]
	j, err := json.Marshal(randomFrasi)
	if err != nil {
		http.Error(w, "non riesco a recuperare le frasi", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, bytes.NewReader(j))
}
