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

// getPokemonsHandler は、すべてのポケモンの情報を取得するエンドポイント
func getPokemonsHandler(w http.ResponseWriter, r *http.Request) {
	// すべてのポケモンを取得するため limit を大きな値に設定
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

	// リスト用レスポンスをパース
	var listResp PokemonListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		http.Error(w, "リストデータの解析に失敗しました", http.StatusInternalServerError)
		return
	}

	// 取得した results の配列を元に、ID・名前・スプライトURLをまとめて返す
	var pokemons []Pokemon

	// ポケモンのURLからIDを取り出すための正規表現
	re := regexp.MustCompile(`/pokemon/(\d+)/?$`)

	for _, result := range listResp.Results {
		matches := re.FindStringSubmatch(result.URL)
		if len(matches) < 2 {
			// ID が取れなかった場合はスキップ
			continue
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			// 数値変換できなければスキップ
			continue
		}

		// スプライト画像は PokeAPI が管理している GitHub から直接参照する形で削減
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

	// 全ポケモン取得エンドポイント
	r.HandleFunc("/pokemons", getPokemonsHandler).Methods("GET")

	fmt.Println("サーバー起動: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
