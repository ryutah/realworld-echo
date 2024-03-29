import { tags } from "@/tests/testdata";
import { Meta, StoryObj } from "@storybook/react";
import PopularTags from "./index";

const meta = {
  title: "index/@pupulartags/PupularTags",
  component: PopularTags,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof PopularTags>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    tags,
  },
};
