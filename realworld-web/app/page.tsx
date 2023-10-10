import { articles, tags } from "@/tests/testdata";
import { Stack } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2/Grid2";
import Main from "./common/layouts/Main";
import FeedTab, { TabType } from "./index/Home/Feed/FeedTab";
import Header from "./index/Home/Header";
import Populartags from "./index/Home/PopularTags/PopularTags";

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
