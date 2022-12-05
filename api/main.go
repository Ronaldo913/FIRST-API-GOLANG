package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/flashcards", getFlashcards)
	fmt.Println("api is on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Flashcard struct {
	Pergunta string `json:"pergunta"`
	Resposta string `json:"resposta"`
}

func getFlashcards(w http.ResponseWriter, request *http.Request) {

	if request.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]Flashcard{
		{
			Pergunta: "Quanto é 1+1?",
			Resposta: "É 2",
		},
		{
			Pergunta: "Quanto é 2+2?",
			Resposta: "É 4",
		},
	})
}
