export interface Pokemon {
  name: string
  url: string
  japanese_name?: string
}

export interface PokemonListResponse {
  results: Pokemon[]
  count: number
} 
