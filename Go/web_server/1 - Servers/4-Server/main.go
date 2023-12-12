package main

import (
	"log"
	"net/http"
)

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
} // funcão copiada do exercício mas deve ser comrpeendida

func main() {
	const port = "8081"

	mux := http.NewServeMux()
	cors := middlewareCors(mux)

	server := http.Server{
		Addr:    ":" + port,
		Handler: cors,
	}
	log.Printf("Servidor rodando na porta: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
