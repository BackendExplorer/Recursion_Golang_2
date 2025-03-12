package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// 単体ポケモン取得エンドポイント
	r.HandleFunc("/pokemon/{id}", getPokemonHandler).Methods("GET")
	// 全ポケモン取得エンドポイント
	r.HandleFunc("/pokemons", getPokemonsHandler).Methods("GET")

	fmt.Println("サーバー起動: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
