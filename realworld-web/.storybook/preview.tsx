// organize-imports-ignore

import React from "react";
import { ThemeProvider } from "@mui/material";
import type { Preview } from "@storybook/react";
import { theme } from "../app/components/ThemeRegistry";

const preview: Preview = {
  parameters: {
    actions: { argTypesRegex: "^on[A-Z].*" },
    nextjs: {
      appDirectory: true,
    },
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/,
      },
    },
  },
};

const withThemeProvider = (Story: Function) => (
  <ThemeProvider theme={theme}>{Story()}</ThemeProvider>
);

export default preview;
export const decorators = [withThemeProvider];
