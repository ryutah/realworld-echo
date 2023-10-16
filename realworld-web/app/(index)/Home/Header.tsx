"use client";

import { Box, Typography, useTheme } from "@mui/material";

export default function Header() {
  const theme = useTheme();

  return (
    <Box
      sx={{
        textAlign: "center",
        background: theme.palette.primary.main,
        width: "100%",
        color: theme.palette.secondary.contrastText,
      }}
      p={5}
    >
      <Typography variant="h3">conduit</Typography>
      <Typography>A place to share your knowledge.</Typography>
    </Box>
  );
}
