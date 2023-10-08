import { LikeButton, TagButton } from "@/app/components/Button";
import { Article, ArticleProps } from "@/app/domain";
import { Avatar, Paper, Stack, Typography } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2/Grid2";

type Props = {
  article: ArticleProps;
};

export default function Feed({ article }: Props) {
  const model = new Article(article);

  return (
    <Paper
      elevation={0}
      sx={{
        width: "100%",
        minWidth: 600,
      }}
    >
      <Grid container spacing={2} p={2}>
        <Grid md={9}>
          <Stack
            direction="row"
            spacing={2}
            sx={{
              alignItems: "center",
            }}
          >
            <Avatar src={model.author.image} alt={article.author.userName} />
            <Stack>
              <Typography>{model.author.userName}</Typography>
              <Typography>{model.formattedUpdatedAt()}</Typography>
            </Stack>
          </Stack>
        </Grid>
        <Grid md={3} sx={{ textAlign: "right", width: "3rem" }}>
          <LikeButton count={article.favoritesCount} />
        </Grid>

        <Grid md={12}>
          <Stack spacing={2}>
            <Stack>
              <Typography variant="h5">{article.title}</Typography>
              <Typography>{article.description}</Typography>
            </Stack>
            <Grid container>
              <Grid md={3}>
                <Typography variant="caption">Read more...</Typography>
              </Grid>
              <Grid md={9} sx={{ textAlign: "right" }}>
                {article.tagList.map((tag) => (
                  <TagButton key={tag}>{tag}</TagButton>
                ))}
              </Grid>
            </Grid>
          </Stack>
        </Grid>
      </Grid>
    </Paper>
  );
}
