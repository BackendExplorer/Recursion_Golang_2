package main

// Sprites はポケモンの画像情報を表します。
type Sprites struct {
	FrontDefault string `json:"front_default"`
}

// PokemonResponse はポケモンAPIのレスポンスを表します。

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

// PokemonDetail はポケモンの詳細情報を表します。
type PokemonDetail struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Image     string   `json:"image"`
	Height    int      `json:"height"`
	Weight    int      `json:"weight"`
	Types     []string `json:"types"`
	Abilities []string `json:"abilities"`
	Stats     Stats    `json:"stats"`
}

// Stats はポケモンのステータス情報を表します。
type Stats struct {
	HP             int `json:"hp"`
	Attack         int `json:"attack"`
	Defense        int `json:"defense"`
	SpecialAttack  int `json:"special_attack"`
	SpecialDefense int `json:"special_defense"`
	Speed          int `json:"speed"`
}

// PokemonResponse はポケモンAPIのレスポンスを表します。
type PokemonResponse struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Height  int     `json:"height"`
	Weight  int     `json:"weight"`
	Sprites Sprites `json:"sprites"`
	Types   []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`
	Stats []struct {
		Stat struct {
			Name string `json:"name"`
		} `json:"stat"`
		BaseStat int `json:"base_stat"`
	} `json:"stats"`
}
