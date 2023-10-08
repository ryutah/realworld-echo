import { Meta, StoryObj } from "@storybook/react";
import { FollowButton } from "./FollowButton";

const meta = {
  title: "Components/Button/FollowButton",
  component: FollowButton,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof FollowButton>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {},
};
