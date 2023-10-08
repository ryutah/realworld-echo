import { Box, Typography } from "@mui/material";

export default function Header() {
  return (
    <Box
      sx={{
        textAlign: "center",
        background: "#5CB85C",
        width: "100%",
        color: "white",
      }}
      p={5}
    >
      <Typography variant="h3">conduit</Typography>
      <Typography>A place to share your knowledge.</Typography>
    </Box>
  );
}
