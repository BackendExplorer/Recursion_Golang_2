export interface Pokemon {
  id: number
  name: string
  image: string
}

export interface PokemonListResponse {
  count: number
  results: Pokemon[]
} 
