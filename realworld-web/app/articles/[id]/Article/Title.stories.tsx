import { articles } from "@/tests/testdata";
import { Meta, StoryObj } from "@storybook/react";
import Title from "./Title";

const meta = {
  title: "article/Article/Title",
  component: Title,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof Title>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    article: articles[0],
  },
};
