import { Meta, StoryObj } from "@storybook/react";
import { TagButton } from "./TagButton";

const meta = {
  title: "common/Components/Button/TagButton",
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
    label: "this is sample tag",
  },
};
