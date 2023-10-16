import { tags } from "@/tests/testdata";
import { Stack } from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2/Grid2";
import Main from "@/app/components/Main";
import FeedTab from "./Home/Feed/FeedTab";
import Header from "./Home/Header";
import Populartags from "./Home/PopularTags/PopularTags";

export default function Home() {
  return (
    <Stack>
      <Header />
      <Main>
        <Grid container>
          <Grid md={9}>
            <FeedTab initTab="global" />
          </Grid>
          <Grid md={3} sx={{ alignItems: "right" }}>
            <Populartags tags={tags} />
          </Grid>
        </Grid>
      </Main>
    </Stack>
  );
}
