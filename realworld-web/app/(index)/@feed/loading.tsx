import { Box, Skeleton, Stack } from "@mui/material";

export default function Loading() {
  return (
    <Stack width={600} alignItems="center">
      <Box>
        <Skeleton variant="circular" width={40} height={40} />
        <Skeleton variant="rectangular" width={210} height={60} />
        <Skeleton variant="rounded" width={210} height={60} />
      </Box>
    </Stack>
  );
}
