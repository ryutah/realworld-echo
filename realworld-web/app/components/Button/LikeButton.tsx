import { Favorite } from "@mui/icons-material";
import { ButtonProps } from "@mui/material";
import NarrowPaddingButton from "./NarrowPaddingButton";

type LikeButtonType = "short" | "long";

type Props = ButtonProps & { liketype?: LikeButtonType; count: number };

const Message = (props: Props) => {
  switch (props.liketype) {
    case "long":
      return ` Favorite Article (${props.count})`;
    default:
      return ` ${props.count}`;
  }
};

export const LikeButton = (props: Props) => (
  <NarrowPaddingButton variant="outlined" color="secondary" {...props}>
    <Favorite fontSize="small" />
    <Message {...props} />
  </NarrowPaddingButton>
);
