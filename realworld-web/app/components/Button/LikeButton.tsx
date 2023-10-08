import { Favorite } from "@mui/icons-material";
import { Button, ButtonProps } from "@mui/material";

type Props = ButtonProps & { count: number };

export const LikeButton = (props: Props) => (
  <Button variant="outlined" {...props}>
    <Favorite fontSize="small" />
    {props.count}
  </Button>
);
