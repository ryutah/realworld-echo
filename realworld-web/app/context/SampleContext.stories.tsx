import { Meta, StoryObj } from "@storybook/react";
import Sample from "./SampleContext";

const meta = {
  title: "Sample/Sample",
  component: Sample,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof Sample>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {};
