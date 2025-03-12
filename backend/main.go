package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Sprites はポケモンの画像情報を表します。
type Sprites struct {
	FrontDefault string `json:"front_default"`
}

// PokemonResponse はポケモンAPIのレスポンスを表します。
type PokemonResponse struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Sprites Sprites `json:"sprites"`
}

// Pokemon はフロントエンド向けのデータ構造
type Pokemon struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

// getPokemonHandler は指定したポケモンの情報を取得するエンドポイント
func getPokemonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// PokeAPI からデータを取得するためのURLを作成して変数に格納する
	apiURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)
	// レスポンスとエラーメッセージを受け取る
	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "外部APIからのデータ取得に失敗しました", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "ポケモンが見つかりません", http.StatusNotFound)
		return
	}

	// レスポンスをパース
	var pokeResponse PokemonResponse
	if err := json.NewDecoder(resp.Body).Decode(&pokeResponse); err != nil {
		http.Error(w, "データの解析に失敗しました", http.StatusInternalServerError)
		return
	}

	// フロントエンド向けのデータ整形
	pokemon := Pokemon{
		ID:    pokeResponse.ID,
		Name:  pokeResponse.Name,
		Image: pokeResponse.Sprites.FrontDefault,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemon)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/pokemon/{id}", getPokemonHandler).Methods("GET")

	fmt.Println("サーバー起動: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
