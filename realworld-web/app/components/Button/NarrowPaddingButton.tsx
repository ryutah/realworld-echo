import { Button, ButtonProps } from "@mui/material";

export default function NarrowPaddingButton(props: ButtonProps) {
  const { sx: _, ...otherSx } = props;
  return (
    <Button
      variant="outlined"
      sx={{
        paddingTop: 0.2,
        paddingBottom: 0.2,
        paddingRight: 0.8,
        paddingLeft: 0.8,
        textTransform: "none",
        ...props.sx,
      }}
      {...otherSx}
    />
  );
}
