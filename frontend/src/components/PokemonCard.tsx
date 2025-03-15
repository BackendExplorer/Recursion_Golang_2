import { Card, CardContent, Typography, Box } from '@mui/material'
import { Pokemon } from '../types/pokemon'

interface Props {
  pokemon: Pokemon
}

export const PokemonCard = ({ pokemon }: Props) => {
  return (
    <Card>
      <CardContent>
        <Typography>
          No.{pokemon.id}
        </Typography>
        <Typography variant="h6" component="h2" sx={{ textTransform: 'capitalize' }}>
          {pokemon.name}
        </Typography>
        <Box sx={{ textAlign: 'center', mt: 2 }}>
          <img 
            src={pokemon.image}
            alt={pokemon.name}
            style={{ width: '80%', height: 'auto' }}
          />
        </Box>
      </CardContent>
    </Card>
  )
} 
