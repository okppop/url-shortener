import {
    AppBar,
    Box,
    Typography,
    Toolbar,
    IconButton,
} from "@mui/material";
import LightModeIcon from '@mui/icons-material/LightMode';
import DarkModeIcon from '@mui/icons-material/DarkMode';

function Bar({ mode, changeMode }) {
    return (
        <Box sx={{ flexGrow: 1 }}>
            <AppBar position="static">
                <Toolbar>
                    <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                        url-shortener
                    </Typography>
                    <IconButton
                        onClick={changeMode}
                    >
                        {mode === "dark" ? <LightModeIcon></LightModeIcon> : <DarkModeIcon></DarkModeIcon>}
                    </IconButton>
                </Toolbar>
            </AppBar>
        </Box>
    )
}

export default Bar