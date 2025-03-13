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

	// 単体ポケモン取得エンドポイント
	r.HandleFunc("/pokemon/{id}", getPokemonHandler).Methods("GET")
	// 全ポケモン取得エンドポイント
	r.HandleFunc("/pokemons", getPokemonsHandler).Methods("GET")

	// CORSミドルウェアの設定
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // フロントエンドのオリジン
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
	})

	// CORSミドルウェアを適用したハンドラを作成
	handler := c.Handler(r)

	fmt.Println("サーバー起動: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
