import { articles } from "@/tests/testdata";
import { Meta, StoryObj } from "@storybook/react";
import FeedList from "./FeedList";

const meta = {
  title: "index/@feed/FeedTab/FeedList",
  component: FeedList,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof FeedList>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    articles: articles,
  },
};
