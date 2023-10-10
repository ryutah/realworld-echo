import Home from "@/app/page";
import { expect } from "@storybook/jest";
import { Meta, StoryObj } from "@storybook/react";
import { within } from "@storybook/testing-library";
import { TestIds as FeedListTestIds } from "./index/Home/Feed/FeedList";

const meta = {
  title: "index/page",
  component: Home,
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
} satisfies Meta<typeof Home>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  play: async ({ canvasElement, step }) => {
    const canvas = within(canvasElement);

    await step("sample", async () => {
      await expect(canvas.getByTestId(FeedListTestIds.List)).toBeTruthy();
    });
  },
};
