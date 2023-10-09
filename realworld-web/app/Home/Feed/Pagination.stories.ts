import { Meta, StoryObj } from "@storybook/react";
import Pagination from "./Pagination";

const meta = {
  title: "Home/Feed/Pagination",
  component: Pagination,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof Pagination>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    count: 500,
  },
};

export const SpecifyPage: Story = {
  parameters: {
    nextjs: {
      navigation: {
        query: {
          page: 201,
        },
      },
    },
  },
  args: {
    count: 500,
  },
};
