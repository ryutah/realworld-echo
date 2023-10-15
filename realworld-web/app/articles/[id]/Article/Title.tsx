import Main from "@/app/components/Main";
import { ArticleProps } from "@/app/domain";
import { Box, Stack, Typography } from "@mui/material";
import Meta from "./Meta";

type Props = {
  article: ArticleProps;
};

export default function Title({ article }: Props) {
  return (
    <Box
      sx={{
        paddingTop: 5,
        paddingBottom: 5,
        background: "#111",
        color: "white",
        width: "100%",
      }}
    >
      <Main>
        <Stack spacing={2}>
          <Typography variant="h3">{article.title}</Typography>
          <Meta article={article} />
        </Stack>
      </Main>
    </Box>
  );
}
