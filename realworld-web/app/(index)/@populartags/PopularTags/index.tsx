import { TagButton } from "@/app/components/Button";
import { Paper, Stack, Typography } from "@mui/material";

type Props = {
  tags: string[];
};

export default function Populartags({ tags }: Props) {
  return (
    <Paper
      elevation={0}
      sx={{
        backgroundColor: "whitesmoke",
        maxWidth: 240,
      }}
    >
      <Stack p={2} spacing={1}>
        <Typography variant="subtitle1">Popular Tags</Typography>
        <Stack direction="row" spacing={0.2} useFlexGap flexWrap="wrap">
          {tags.map((tag) => (
            <TagButton key={tag} label={tag} clickable={true} />
          ))}
        </Stack>
      </Stack>
    </Paper>
  );
}
