package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const pokeAPIBaseURL = "https://pokeapi.co/api/v2"

// 単体ポケモン取得処理
func getPokemonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	res, err := http.Get(fmt.Sprintf("%s/pokemon/%s", pokeAPIBaseURL, id))
	if err != nil {
		http.Error(w, "ポケモン情報取得エラー", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	var pokeResp PokemonResponse
	json.NewDecoder(res.Body).Decode(&pokeResp)

	pokemon := Pokemon{
		ID:    pokeResp.ID,
		Name:  pokeResp.Name,
		Image: pokeResp.Sprites.FrontDefault,
	}

	json.NewEncoder(w).Encode(pokemon)
}

// 全ポケモン取得処理（ページネーション機能付き）
func getPokemonsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// ページネーションのパラメータ取得（デフォルト値を設定）
	page, _ := strconv.Atoi(query.Get("page"))
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(query.Get("limit"))
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	offset := (page - 1) * limit

	url := fmt.Sprintf("%s/pokemon?offset=%d&limit=%d", pokeAPIBaseURL, offset, limit)

	res, err := http.Get(url)
	if err != nil {
		http.Error(w, "ポケモン情報一覧取得エラー", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	var listResp PokemonListResponse
	json.NewDecoder(res.Body).Decode(&listResp)

	// 各ポケモンの詳細を取得し、必要な情報を配列化
	pokemons := make([]Pokemon, 0, len(listResp.Results))

	for _, result := range listResp.Results {
		pokemonRes, err := http.Get(result.URL)
		if err != nil {
			continue // エラーがあれば飛ばす
		}

		var pokeResp PokemonResponse
		json.NewDecoder(pokemonRes.Body).Decode(&pokeResp)
		pokemonRes.Body.Close()

		pokemons = append(pokemons, Pokemon{
			ID:    pokeResp.ID,
			Name:  pokeResp.Name,
			Image: pokeResp.Sprites.FrontDefault,
		})
	}

	json.NewEncoder(w).Encode(struct {
		Page     int       `json:"page"`
		Limit    int       `json:"limit"`
		Pokemons []Pokemon `json:"pokemons"`
	}{
		Page:     page,
		Limit:    limit,
		Pokemons: pokemons,
	})
}
