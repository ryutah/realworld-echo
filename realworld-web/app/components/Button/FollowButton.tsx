import { Add } from "@mui/icons-material";
import { ButtonProps } from "@mui/material";
import NarrowPaddingButton from "./NarrowPaddingButton";

type Props = ButtonProps & { user: string };

export const FollowButton = (props: Props) => (
  <NarrowPaddingButton variant="outlined" color="info" {...props}>
    <Add fontSize="small" />
    {` Follow ${props.user}`}
  </NarrowPaddingButton>
);
