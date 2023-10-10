import { articles } from "@/tests/testdata";
import { Meta, StoryObj } from "@storybook/react";
import FeedTab, { TabType } from "./FeedTab";

const meta = {
  title: "index/Home/Feed/FeedTab",
  component: FeedTab,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof FeedTab>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    articles: articles,
    initTab: TabType.Global,
  },
};
