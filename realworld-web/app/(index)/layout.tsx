import { Stack } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2/Grid2";
import { ReactNode } from "react";
import Main from "../components/Main";

type Props = {
  children: ReactNode;
  feed: ReactNode;
  populartags: ReactNode;
};

export default function ArticleLayout(props: Props) {
  return (
    <Stack>
      {props.children}
      <Main>
        <Grid container>
          <Grid md={9}>{props.feed}</Grid>
          <Grid md={3} sx={{ alignItems: "right" }}>
            {props.populartags}
          </Grid>
        </Grid>
      </Main>
    </Stack>
  );
}
