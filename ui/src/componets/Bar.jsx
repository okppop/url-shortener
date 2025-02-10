import {
    AppBar,
    Box,
    Typography,
    Toolbar,
} from "@mui/material";

function Bar() {
    return (
        <Box sx={{ flexGrow: 1 }}>
            <AppBar position="static">
                <Toolbar>
                    <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                        url-shortener
                    </Typography>
                </Toolbar>
            </AppBar>
        </Box>
    )
}

export default Bar