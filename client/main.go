package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Flashcard struct {
	Pergunta string `json:"pergunta"`
	Resposta string `json:"resposta"`
}

func main() {
  resp, err := http.Get("http://localhost:8080/flashcards")
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  if resp.StatusCode != 200 {
    fmt.Println("Not sucess", resp.StatusCode)
    return
  }

  body, err := io.ReadAll(resp.Body)

  var response []Flashcard
  err = json.Unmarshal(body, &response)
  if err != nil {
    fmt.Println("Erro ao recuperar dados", err.Error())
    return
  }

  fmt.Println(response)
}