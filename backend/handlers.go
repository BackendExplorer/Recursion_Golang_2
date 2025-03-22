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
// ブラウザからの要求を受け取り、ブラウザに要求されたものを返す
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

	pokemonDetail := PokemonDetail{
		ID:        pokeResp.ID,
		Name:      pokeResp.Name,
		Image:     pokeResp.Sprites.FrontDefault,
		Height:    pokeResp.Height,
		Weight:    pokeResp.Weight,
		Types:     make([]string, len(pokeResp.Types)),
		Abilities: make([]string, len(pokeResp.Abilities)),
		Stats:     Stats{},
	}

	for i, t := range pokeResp.Types {
		pokemonDetail.Types[i] = t.Type.Name
	}

	for i, a := range pokeResp.Abilities {
		pokemonDetail.Abilities[i] = a.Ability.Name
	}

	for _, s := range pokeResp.Stats {
		switch s.Stat.Name {
		case "hp":
			pokemonDetail.Stats.HP = s.BaseStat
		case "attack":
			pokemonDetail.Stats.Attack = s.BaseStat
		case "defense":
			pokemonDetail.Stats.Defense = s.BaseStat
		case "special-attack":
			pokemonDetail.Stats.SpecialAttack = s.BaseStat
		case "special-defense":
			pokemonDetail.Stats.SpecialDefense = s.BaseStat
		case "speed":
			pokemonDetail.Stats.Speed = s.BaseStat
		}
	}

	if err := json.NewEncoder(w).Encode(pokemonDetail); err != nil {
		http.Error(w, "JSONエンコードエラー", http.StatusInternalServerError)
		return
	}
}

// 全ポケモン取得処理（ページネーション対応）
func getPokemonsHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータからページとリミットを取得
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "20"
	}

	// offsetを計算（ページ - 1）* リミット
	offset := 0
	if pageNum, err := strconv.Atoi(page); err == nil && pageNum > 1 {
		limitNum, _ := strconv.Atoi(limit)
		offset = (pageNum - 1) * limitNum
	}

	url := fmt.Sprintf("%s/pokemon?offset=%d&limit=%s", pokeAPIBaseURL, offset, limit)

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

		// 即座にCloseするように修正
		var pokeResp PokemonResponse
		if err := json.NewDecoder(pokemonRes.Body).Decode(&pokeResp); err != nil {
			pokemonRes.Body.Close()
			continue
		}
		pokemonRes.Body.Close()

		pokemons = append(pokemons, Pokemon{
			ID:    pokeResp.ID,
			Name:  pokeResp.Name,
			Image: pokeResp.Sprites.FrontDefault,
		})
	}

	// ページネーション情報を含めたレスポンスを返す
	if err := json.NewEncoder(w).Encode(struct {
		Pokemon []Pokemon `json:"pokemon"`
		Total   int       `json:"total"`
		Page    string    `json:"page"`
		Limit   string    `json:"limit"`
	}{
		Pokemon: pokemons,
		Total:   listResp.Count,
		Page:    page,
		Limit:   limit,
	}); err != nil {
		http.Error(w, "JSONエンコードエラー", http.StatusInternalServerError)
		return
	}
}
