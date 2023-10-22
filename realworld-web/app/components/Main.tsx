import { Box, BoxProps } from "@mui/material";

type Props = BoxProps;

export default function Main(props: Props) {
  const { sx: _, ...otherSx } = props;

  return (
    <Box
      sx={{
        ...props.sx,
        maxWidth: "1080px",
        margin: "0 auto",
        padding: 2,
      }}
      {...otherSx}
    />
  );
}
