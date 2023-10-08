import { Meta, StoryObj } from "@storybook/react";
import { TagButton } from "./TagButton";

const meta = {
  title: "Components/Button/TagButton",
  component: TagButton,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof TagButton>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    children: "this is sample tag",
  },
};
