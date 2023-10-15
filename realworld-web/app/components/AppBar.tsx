import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";

export default function MyAppBar() {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            conduit
          </Typography>
          <Button color="inherit">Home</Button>
          <Button color="inherit">Sign in</Button>
          <Button color="inherit">Sign up</Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
