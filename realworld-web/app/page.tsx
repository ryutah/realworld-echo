import Header from "@/app/Home/Header";
import Grid from "@mui/material/Unstable_Grid2/Grid2";
import Populartags from "./Home/PopularTags/PopularTags";

import { articles, tags } from "@/tests/testdata";
import FeedTab, { TabType } from "./Home/Feed/FeedTab";

export default function Home() {
  return (
    <Grid container spacing={2} sx={{ minWidth: 500 }}>
      <Grid md={12}>
        <Header />
      </Grid>
      <Grid md={9}>
        <FeedTab initTab={TabType.Global} articles={articles} />
      </Grid>
      <Grid md={3} sx={{ alignItems: "right" }}>
        <Populartags tags={tags} />
      </Grid>
    </Grid>
  );
}
