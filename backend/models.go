package main

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
