import { Meta, StoryObj } from "@storybook/react";
import MyAppBar from "./AppBar";

const meta = {
  title: "Components/AppBar",
  component: MyAppBar,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof MyAppBar>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {};
