import { ArticleProps } from "@/app/domain";
import { Box, Divider, Stack } from "@mui/material";
import Feed from "./Feed";

export const TestIds = {
  List: "home/feed-list",
};

type Props = {
  articles: ArticleProps[];
};

export default function FeedList({ articles }: Props) {
  return (
    <Box data-testid={TestIds.List}>
      <Stack divider={<Divider orientation="horizontal" flexItem />}>
        {articles.map((a) => (
          <Feed key={a.slug} article={a} />
        ))}
      </Stack>
    </Box>
  );
}
