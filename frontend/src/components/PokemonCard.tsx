import { Card, CardContent, Typography, Box } from '@mui/material'
import { Pokemon } from '../types/pokemon'

interface Props {
  pokemon: Pokemon
}

export const PokemonCard = ({ pokemon }: Props) => {
  const id = pokemon.url.split('/').filter(Boolean).pop()

  return (
    <Card>
      <CardContent>
        <Typography>
          No.{id}
        </Typography>
        <Typography variant="h6" component="h2" sx={{ textTransform: 'capitalize' }}>
          {pokemon.japanese_name || pokemon.name}
        </Typography>
        <Box sx={{ textAlign: 'center', mt: 2 }}>
          <img 
            src={`https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/${id}.png`}
            alt={pokemon.japanese_name || pokemon.name}
            style={{ width: '80%', height: 'auto' }}
          />
        </Box>
      </CardContent>
    </Card>
  )
} 
