import { PokemonListResponse, Pokemon } from '../types/pokemon'

export const fetchPokemonList = async (page: number): Promise<PokemonListResponse> => {
  const offset = (page - 1) * 20
  const response = await fetch(`https://pokeapi.co/api/v2/pokemon-species?offset=${offset}&limit=20`)
  
  if (!response.ok) {
    throw new Error('ネットワークエラーが発生しました')
  }
  
  const data = await response.json()
  const pokemonWithJapaneseNames = await Promise.all(
    data.results.map(async (pokemon: Pokemon) => {
      const speciesResponse = await fetch(pokemon.url)
      const speciesData = await speciesResponse.json()
      
      const japaneseName = speciesData.names.find(
        (name: any) => name.language.name === 'ja'
      )?.name || pokemon.name
      
      return {
        ...pokemon,
        japanese_name: japaneseName
      }
    })
  )
  
  return {
    ...data,
    results: pokemonWithJapaneseNames
  }
} 
