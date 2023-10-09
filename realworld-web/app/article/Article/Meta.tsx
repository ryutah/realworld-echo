import { FollowButton, LikeButton } from "@/app/common/components/Button";
import { Article, ArticleProps } from "@/app/domain";
import { Avatar, Stack, Typography } from "@mui/material";

type Props = {
  article: ArticleProps;
};

export default function Meta({ article }: Props) {
  const model = new Article(article);
  return (
    <Stack direction="row" spacing={1} alignItems="center">
      <Avatar src={model.author.image} />
      <Stack>
        <Typography variant="subtitle1">{model.author.userName}</Typography>
        <Typography variant="caption">{model.formattedUpdatedAt()}</Typography>
      </Stack>
      <FollowButton user={article.author.userName} />
      <LikeButton liketype="long" count={article.favoritesCount} />
    </Stack>
  );
}
