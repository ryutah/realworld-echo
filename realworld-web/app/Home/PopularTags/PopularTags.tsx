import { TagButton } from "@/app/components/Button";
import { Paper, Typography } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2/Grid2";

type Props = {
  tags: string[];
};

export default function Populartags({ tags }: Props) {
  return (
    <Paper
      elevation={0}
      sx={{
        backgroundColor: "whitesmoke",
        maxWidth: "15rem",
      }}
    >
      <Grid container p={2} spacing={1}>
        <Grid md={12}>
          <Typography variant="subtitle1">Popular Tags</Typography>
        </Grid>
        <Grid md={12} spacing={2}>
          {tags.map((tag) => (
            <TagButton key={tag} label={tag} sx={{ margin: 0.1 }} />
          ))}
        </Grid>
      </Grid>
    </Paper>
  );
}
