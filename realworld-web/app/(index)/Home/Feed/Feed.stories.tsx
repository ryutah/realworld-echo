import { articles } from "@/tests/testdata";
import { Meta, StoryObj } from "@storybook/react";
import Feed from "./Feed";

const meta = {
  title: "index/Home/Feed/Feed",
  component: Feed,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  decorators: [
    (Story) => (
      <div>
        <Story />
      </div>
    ),
  ],
  argTypes: {},
} satisfies Meta<typeof Feed>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    article: articles[0],
  },
};
