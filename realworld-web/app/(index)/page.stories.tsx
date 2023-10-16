import { articles } from "@/tests/testdata";
import { expect } from "@storybook/jest";
import { Meta, StoryObj } from "@storybook/react";
import { within } from "@storybook/testing-library";
import { TestIds as FeedListTestIds } from "./Home/Feed/FeedList";
import Page from "./page";
import { ArticlesProvider } from "./store";

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
  decorators: [
    (Stroy) => (
      <ArticlesProvider globalFeeds={articles}>
        <Stroy />
      </ArticlesProvider>
    ),
  ],
} satisfies Meta<typeof Page>;

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
