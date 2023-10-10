import { Meta, StoryObj } from "@storybook/react";
import NarrowPaddingButton from "./NarrowPaddingButton";

const meta = {
  title: "common/components/Button/NarrowPaddingButton",
  component: NarrowPaddingButton,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof NarrowPaddingButton>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    children: "Button Sample",
  },
};
