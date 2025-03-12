import { useQuery } from '@tanstack/react-query'
import { 
  Container, 
  Typography, 
  Box, 
  CircularProgress, 
  Grid as Unstable_Grid2,
  AppBar,
  Toolbar,
  Pagination,
} from '@mui/material'
import React from 'react'
import { PokemonCard } from './components/PokemonCard'
import { fetchPokemonList } from './api/pokemon'
import { PokemonListResponse } from './types/pokemon'

function App() {
  const [page, setPage] = React.useState(1)
  const itemsPerPage = 20

  const { data, isLoading, error } = useQuery<PokemonListResponse, Error>({
    queryKey: ['pokemon', page],
    queryFn: () => fetchPokemonList(page),
  })

  const handlePageChange = (_event: React.ChangeEvent<unknown>, value: number) => {
    setPage(value)
  }

  return (
    <>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            ポケモン図鑑
          </Typography>
        </Toolbar>
      </AppBar>
      
      <Container maxWidth="lg" sx={{ mt: 4 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          ポケモン一覧
        </Typography>
        
        {isLoading ? (
          <Box sx={{ display: 'flex', justifyContent: 'center', mt: 4 }}>
            <CircularProgress />
          </Box>
        ) : error ? (
          <Typography color="error">エラー: {error.message}</Typography>
        ) : (
          <Unstable_Grid2 container spacing={2}>
            {data?.results.map((pokemon) => (
              <Unstable_Grid2 xs={5} sm={3} key={pokemon.name}>
                <PokemonCard pokemon={pokemon} />
              </Unstable_Grid2>
            ))}
          </Unstable_Grid2>
        )}
        
        {!isLoading && !error && data && (
          <Box sx={{ display: 'flex', justifyContent: 'center', mt: 4, mb: 4 }}>
            <Pagination 
              count={Math.ceil((data.count || 0) / itemsPerPage)}
              page={page}
              onChange={handlePageChange}
              color="primary"
              size="large"
            />
          </Box>
        )}
      </Container>
    </>
  )
}

export default App
