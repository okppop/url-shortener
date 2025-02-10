import {
    Accordion,
    AccordionDetails,
    AccordionSummary,
    Box,
    Paper,
    Typography,
} from "@mui/material"
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import URL from "./URL"

function Info() {
    return (
        <Paper square={true}>
            <Box sx={{
                paddingTop: "50px",
                paddingLeft: "5%",
                paddingRight: "5%",
                display: "flex",
                flexDirection: "column",
                gap: 5,
                // alignItems: "center",
            }}>
                <Typography variant="h2">
                    url-shortener
                </Typography>
                <Typography variant="h6">
                    This service intend to make any URL short.
                </Typography>
                <Accordion sx={{}}>
                    <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                        <Typography variant="h6" component="span">
                            What's that means?
                        </Typography>
                    </AccordionSummary>

                    <AccordionDetails>
                        <Typography variant="h6" component="p">
                            For example, you have a long URL like this:
                        </Typography>
                        <Typography variant="p" >
                            https://en.wikipedia.org/wiki/Special:BookSources/978-0-13-854662-5
                        </Typography>
                        <Box sx={{ padding: "1%" }}></Box>
                        <Typography variant="h6" component="p">
                            You can make it like this:
                        </Typography>
                        <Typography variant="p">
                            {window.location.origin + "/Iy6Ngx"}
                        </Typography>
                    </AccordionDetails>
                </Accordion>
                <URL />
                <Box sx={{
                    marginBottom: "50%",
                }}>
                </Box>
            </Box>
        </Paper>
    )
}

export default Info