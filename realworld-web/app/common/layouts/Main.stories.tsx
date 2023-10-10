import { Meta, StoryObj } from "@storybook/react";
import Main from "./Main";

const meta = {
  title: "common/layouts/Main",
  component: Main,
  parameters: {
    layout: "centered",
  },
  decorators: [
    (Story) => (
      <div style={{ background: "red", height: "200px", width: "1400px" }}>
        <Story />
      </div>
    ),
  ],
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof Main>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    children: "this is main contents",
    sx: {
      background: "blue",
    },
  },
};
