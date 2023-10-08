import Home from "@/app/page";
import { Meta, StoryObj } from "@storybook/react";

const meta = {
  title: "Home/page",
  component: Home,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof Home>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {};
