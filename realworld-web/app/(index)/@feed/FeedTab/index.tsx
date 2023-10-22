"use client";

import { Box, Stack, Tab, Tabs } from "@mui/material";
import { useState } from "react";
import { useArticles } from "../store";
import FeedList from "./FeedList";
import Pagination from "./Pagination";

export const TestIds = {
  Tab: "home/feed-tab",
  GlobalTab: "home/feed-tab/global-tab",
};

const TabType = {
  Global: "global",
} as const;
export type TabType = (typeof TabType)[keyof typeof TabType];

type Props = {
  initTab: TabType;
};

export default function FeedTab({ initTab }: Props) {
  const store = useArticles();
  const [currentTab, setCurrentTab] = useState(initTab);

  return (
    <Box data-testid={TestIds.Tab}>
      <Stack alignItems="center">
        <Tabs
          value={currentTab}
          onChange={(_, newValue) => setCurrentTab(newValue)}
        >
          <Tab
            value={TabType.Global}
            sx={{
              textTransform: "none",
            }}
            label="Global Feed"
          />
        </Tabs>

        <Box
          data-testid={TestIds.GlobalTab}
          hidden={currentTab !== TabType.Global}
        >
          <FeedList articles={store.articles.globalFeeds} />
        </Box>
        <Box>
          <Pagination count={store.articles.globalFeeds.length} />
        </Box>
      </Stack>
    </Box>
  );
}
