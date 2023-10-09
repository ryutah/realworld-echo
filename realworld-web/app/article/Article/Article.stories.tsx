import { articles } from "@/tests/testdata";
import { Meta, StoryObj } from "@storybook/react";
import Article from "./Article";

const meta = {
  title: "article/Article/Article",
  component: Article,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof Article>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    article: articles[0],
  },
};
