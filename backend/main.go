package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

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

// Pokemon はフロントエンド向けのデータ構造（単体用・複数用で共通利用）。
type Pokemon struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

// PokemonListResponse は /pokemon?limit=... 取得時のレスポンスを表します。
type PokemonListResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// getPokemonHandler は指定したポケモンの情報を取得するエンドポイント
func getPokemonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// PokeAPI からデータを取得するためのURLを作成
	apiURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)
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

	var pokeResponse PokemonResponse
	if err := json.NewDecoder(resp.Body).Decode(&pokeResponse); err != nil {
		http.Error(w, "データの解析に失敗しました", http.StatusInternalServerError)
		return
	}

	pokemon := Pokemon{
		ID:    pokeResponse.ID,
		Name:  pokeResponse.Name,
		Image: pokeResponse.Sprites.FrontDefault,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemon)
}

// getPokemonsHandler は、すべてのポケモンの情報を取得するエンドポイント
func getPokemonsHandler(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://pokeapi.co/api/v2/pokemon?limit=2000"
	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "外部APIからのリスト取得に失敗しました", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "ポケモンリストが見つかりません", http.StatusNotFound)
		return
	}

	var listResp PokemonListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		http.Error(w, "リストデータの解析に失敗しました", http.StatusInternalServerError)
		return
	}

	var pokemons []Pokemon
	re := regexp.MustCompile(`/pokemon/(\d+)/?$`)

	for _, result := range listResp.Results {
		matches := re.FindStringSubmatch(result.URL)
		if len(matches) < 2 {
			continue
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			continue
		}

		imageURL := fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/%d.png", id)
		pokemons = append(pokemons, Pokemon{
			ID:    id,
			Name:  result.Name,
			Image: imageURL,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func main() {
	r := mux.NewRouter()

	// 単体ポケモン取得エンドポイント
	r.HandleFunc("/pokemon/{id}", getPokemonHandler).Methods("GET")
	// 全ポケモン取得エンドポイント
	r.HandleFunc("/pokemons", getPokemonsHandler).Methods("GET")

	fmt.Println("サーバー起動: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
