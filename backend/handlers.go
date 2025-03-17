package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	if err := json.NewDecoder(res.Body).Decode(&pokeResp); err != nil {
		http.Error(w, "JSONデコードエラー", http.StatusInternalServerError)
		return
	}

	pokemon := Pokemon{
		ID:    pokeResp.ID,
		Name:  pokeResp.Name,
		Image: pokeResp.Sprites.FrontDefault,
	}

	if err := json.NewEncoder(w).Encode(pokemon); err != nil {
		http.Error(w, "JSONエンコードエラー", http.StatusInternalServerError)
		return
	}
}

// 全ポケモン取得処理（page / limit を削除）
func getPokemonsHandler(w http.ResponseWriter, r *http.Request) {
	// ここでは例として20件を固定で取得する
	url := fmt.Sprintf("%s/pokemon?limit=20", pokeAPIBaseURL)

	res, err := http.Get(url)
	if err != nil {
		http.Error(w, "ポケモン情報一覧取得エラー", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	var listResp PokemonListResponse
	if err := json.NewDecoder(res.Body).Decode(&listResp); err != nil {
		http.Error(w, "JSONデコードエラー", http.StatusInternalServerError)
		return
	}

	pokemons := make([]Pokemon, 0, len(listResp.Results))

	for _, result := range listResp.Results {
		pokemonRes, err := http.Get(result.URL)
		if err != nil {
			continue
		}
		defer pokemonRes.Body.Close()

		var pokeResp PokemonResponse
		if err := json.NewDecoder(pokemonRes.Body).Decode(&pokeResp); err != nil {
			continue
		}

		pokemons = append(pokemons, Pokemon{
			ID:    pokeResp.ID,
			Name:  pokeResp.Name,
			Image: pokeResp.Sprites.FrontDefault,
		})
	}

	// page / limit のフィールドは削除し、ポケモン配列のみ返す
	if err := json.NewEncoder(w).Encode(struct {
		Pokemon []Pokemon `json:"pokemon"`
	}{
		Pokemon: pokemons,
	}); err != nil {
		http.Error(w, "JSONエンコードエラー", http.StatusInternalServerError)
		return
	}
}
