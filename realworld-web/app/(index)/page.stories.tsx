import { Meta, StoryObj } from "@storybook/react";
import Page from "./page";

const meta = {
  title: "index/page",
  component: Page,
  parameters: {
    layout: "centered",
    nextjs: {
      navigation: {
        query: {
          page: 201,
        },
      },
    },
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof Page>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {};
