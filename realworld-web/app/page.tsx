import Header from "@/app/Home/Header";
import Grid from "@mui/material/Unstable_Grid2/Grid2";
import Populartags from "./Home/PopularTags/PopularTags";

import { articles, tags } from "@/tests/testdata";
import { Stack } from "@mui/material";
import FeedTab, { TabType } from "./Home/Feed/FeedTab";
import Main from "./common/layouts/Main";

export default function Home() {
  return (
    <Stack>
      <Header />
      <Main>
        <Grid container>
          <Grid md={9}>
            <FeedTab initTab={TabType.Global} articles={articles} />
          </Grid>
          <Grid md={3} sx={{ alignItems: "right" }}>
            <Populartags tags={tags} />
          </Grid>
        </Grid>
      </Main>
    </Stack>
  );
}
