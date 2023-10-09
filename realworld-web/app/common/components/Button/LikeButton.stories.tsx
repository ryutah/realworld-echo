import { Meta, StoryObj } from "@storybook/react";
import { LikeButton } from "./LikeButton";

const meta = {
  title: "common/Components/Button/LikeButton",
  component: LikeButton,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof LikeButton>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    count: 10,
  },
};

export const LongMessage: Story = {
  args: {
    count: 10,
    liketype: "long",
  },
};

export const ShortMessage: Story = {
  args: {
    count: 10,
    liketype: "short",
  },
};
