import { articles } from "@/tests/testdata";
import { Meta, StoryObj } from "@storybook/react";
import { ArticlesProvider } from "../../store";
import FeedTab from "./FeedTab";

const meta = {
  title: "index/Home/Feed/FeedTab",
  component: FeedTab,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  argTypes: {},
  decorators: [
    (Story) => (
      <ArticlesProvider globalFeeds={articles}>
        <Story />
      </ArticlesProvider>
    ),
  ],
} satisfies Meta<typeof FeedTab>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    initTab: "global",
  },
};
