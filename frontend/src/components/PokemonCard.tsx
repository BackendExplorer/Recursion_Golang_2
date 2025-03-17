import { Card, CardContent, Typography, Box } from '@mui/material'
import { Pokemon } from '../types/pokemon'

interface Props {
  pokemon: Pokemon
}

export const PokemonCard = ({ pokemon }: Props) => {
  return (
    <Card sx={{ width: 200, height: 280 }}>
      <CardContent>
        <Typography>
          No.{pokemon.id}
        </Typography>
        <Typography 
          variant="h6" 
          component="h2" 
          sx={{ 
            textTransform: 'capitalize',
            overflow: 'hidden',
            textOverflow: 'ellipsis',
            whiteSpace: 'nowrap'
          }}
        >
          {pokemon.name}
        </Typography>
        <Box sx={{ textAlign: 'center', mt: 2 }}>
          <img 
            src={pokemon.image}
            alt={pokemon.name}
            style={{ width: '120px', height: '120px', objectFit: 'contain' }}
          />
        </Box>
      </CardContent>
    </Card>
  )
} 
