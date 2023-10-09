import { ArticleProps } from "@/app/domain";
import { Box, Stack, Tab, Tabs } from "@mui/material";
import { useState } from "react";
import FeedList from "./FeedList";
import Pagination from "./Pagination";

export const TestIds = {
  Tab: "home/feed-tab",
  GlobalTab: "home/feed-tab/global-tab",
};

export const TabType = {
  Global: "global",
} as const;
export type TabType = (typeof TabType)[keyof typeof TabType];

type Props = {
  initTab: TabType;
  articles: ArticleProps[];
};

export default function FeedTab({ initTab, articles }: Props) {
  const [currentTab, setCurrentTab] = useState(initTab);

  return (
    <Box data-testid={TestIds.Tab}>
      <Stack alignItems="center">
        <Tabs
          value={currentTab}
          onChange={(_, newValue) => setCurrentTab(newValue)}
        >
          <Tab value={TabType.Global} label="Global Feed" />
        </Tabs>

        <Box
          data-testid={TestIds.GlobalTab}
          hidden={currentTab !== TabType.Global}
        >
          <FeedList articles={articles} />
        </Box>
        <Box>
          <Pagination count={articles.length} />
        </Box>
      </Stack>
    </Box>
  );
}
