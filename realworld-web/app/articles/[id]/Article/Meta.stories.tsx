import { articles } from "@/tests/testdata";
import { Meta, StoryObj } from "@storybook/react";
import ArticleMeta from "./Meta";

const meta = {
  title: "article/Article/Meta",
  component: ArticleMeta,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof ArticleMeta>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    article: articles[0],
  },
};
