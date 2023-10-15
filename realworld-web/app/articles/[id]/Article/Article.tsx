import { TagButton } from "@/app/components/Button";
import Main from "@/app/components/Main";
import { Article, ArticleProps } from "@/app/domain";
import { Box, Divider, Stack, Typography } from "@mui/material";
import Meta from "./Meta";
import Title from "./Title";

type Props = {
  article: ArticleProps;
};

export default function Articles({ article }: Props) {
  const model = new Article(article);

  return (
    <Stack spacing={2}>
      <Title article={model} />
      <Main>
        <Typography variant="body1">{model.body}</Typography>
        <Box mt={5}>
          {model.tagList.map((tag) => (
            <TagButton key={tag} label={tag} variant="outlined" />
          ))}
        </Box>
        <Box pt={2} pb={2}>
          <Divider />
        </Box>
        <Stack alignItems="center">
          <Meta article={model} />
        </Stack>
      </Main>
    </Stack>
  );
}
