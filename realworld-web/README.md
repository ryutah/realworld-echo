# realworld-web

## How to init project

```console
pnpm create next-app realworld-web

# mui
pnpm add @mui/material @emotion/react @emotion/styled @mui/icons-material

# storybook
pnpm dlx storybook@latest init
pnpm add -D @storybook/testing-library @storybook/jest @storybook/addon-interactions

# jest
pnpm add -D jest ts-node @types/jest ts-jest @testing-library/jest-dom @testing-library/react jest-environment-jsdom

pnpm typesync
```

## Referenses

- Directory Structure
  - [How to Structure Your Next.js App With the New App Router | by Alen Ajam | Better Programming](https://betterprogramming.pub/how-to-structure-your-next-js-app-with-the-new-app-router-61bf2bf5a20d)
- Testing
  - [ESM Support | ts-jest](https://kulshekhar.github.io/ts-jest/docs/next/guides/esm-support/)
