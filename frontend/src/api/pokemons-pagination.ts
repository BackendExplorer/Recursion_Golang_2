import { PokemonListResponse } from '../types/pokemon'
import { useQuery } from '@tanstack/react-query'

export const fetchPokemonList = async (page: number = 1, limit: number = 20): Promise<PokemonListResponse> => {
  try {
    const response = await fetch(`http://localhost:8080/pokemons?page=${page}&limit=${limit}`, {
      headers: {
        'Accept': 'application/json',
      },
      signal: AbortSignal.timeout(10000),
    })
    
    if (!response.ok) {
      throw new Error(`ネットワークエラーが発生しました: ${response.status} ${response.statusText}`)
    }
    
    const data = await response.json()
    return {
      count: data.total,
      results: data.pokemon
    }
  } catch (error) {
    if (error instanceof TypeError && error.message === 'Failed to fetch') {
      throw new Error('サーバーに接続できませんでした。サーバーが起動しているか確認してください。')
    }
    throw error
  }
}

export const usePokemonList = (page: number = 1, limit: number = 20) => {
  return useQuery({
    queryKey: ['pokemons', page, limit],
    queryFn: () => fetchPokemonList(page, limit),
  })
} 
