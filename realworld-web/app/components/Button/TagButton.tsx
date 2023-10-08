import { Button, ButtonProps } from "@mui/material";

type Props = ButtonProps;

export const TagButton = (props: Props) => (
  <Button
    size="small"
    variant="outlined"
    sx={{
      fontSize: "0.5rem",
      borderRadius: 10,
    }}
    {...props}
  />
);
