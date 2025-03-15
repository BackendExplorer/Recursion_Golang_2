import { PokemonListResponse } from '../types/pokemon'
import { useQuery } from '@tanstack/react-query'

export const fetchPokemonList = async (): Promise<PokemonListResponse> => {
  try {
    const response = await fetch(`http://localhost:8080/pokemons`, {
      headers: {
        'Accept': 'application/json',
      },
      signal: AbortSignal.timeout(10000),
    })
    
    if (!response.ok) {
      throw new Error(`ネットワークエラーが発生しました: ${response.status} ${response.statusText}`)
    }
    
    const pokemons = await response.json()
    return {
      count: pokemons.length,
      results: pokemons
    }
  } catch (error) {
    if (error instanceof TypeError && error.message === 'Failed to fetch') {
      throw new Error('サーバーに接続できませんでした。サーバーが起動しているか確認してください。')
    }
    throw error
  }
}

export const usePokemonList = () => {
  return useQuery({
    queryKey: ['pokemons'],
    queryFn: fetchPokemonList,
  })
} 
