package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	// エンドポイント
	//pokemon/{id}というパスを持つURLが入力された時だけ、getPokemonHandlerという関数を呼び出すということを明示しているのか？
	r.HandleFunc("/pokemon/{id}", getPokemonHandler).Methods("GET")
	r.HandleFunc("/pokemons", getPokemonsHandler).Methods("GET")

	// CORS設定
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
	})

	handler := c.Handler(r)

	fmt.Println("サーバー起動: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
