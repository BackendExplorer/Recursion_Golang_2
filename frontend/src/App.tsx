import { useQuery } from '@tanstack/react-query'
import { 
  Container, 
  Typography, 
  Box, 
  CircularProgress, 
  Card, 
  CardContent,
  Grid,
  AppBar,
  Toolbar,
} from '@mui/material'


// 投稿データの型定義を追加
interface Post {
  id: number;
  title: string;
  body: string;
}

// データフェッチ関数の戻り値の型を指定
const fetchPosts = async (): Promise<Post[]> => {
  const response = await fetch('https://jsonplaceholder.typicode.com/posts')
  if (!response.ok) {
    throw new Error('ネットワークエラーが発生しました')
  }
  return response.json()
}

function App() {
  // useQueryの型パラメータを指定
  const { data, isLoading, error } = useQuery<Post[], Error>({
    queryKey: ['posts'],
    queryFn: fetchPosts,
  })

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
          <Grid container spacing={3}>
            {data?.slice(0, 10).map((post: Post) => (
              <Grid item xs={12} sm={6} md={4} key={post.id}>
                <Card>
                  <CardContent>
                    <Typography variant="h6" component="h2">
                      {post.title}
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                      {post.body}
                    </Typography>
                  </CardContent>
                </Card>
              </Grid>
            ))}
          </Grid>
        )}
      </Container>
    </>
  )
}

export default App
