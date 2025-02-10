import { Alert, Box, Button, FormControl, FormControlLabel, FormGroup, InputLabel, MenuItem, Select, Switch, TextField, Typography } from "@mui/material"
import { useState } from "react"
import { CreateURL } from "../api/api"

function URL() {
    const [originalURL, setOriginalURL] = useState("")
    const [expireHours, setExpireHours] = useState(12)
    const [neverExpire, setNeverExpire] = useState(false)
    const changeOriginalURL = (e) => {
        setOriginalURL(e.target.value)
    }
    const changeExpireHours = (e) => {
        setExpireHours(e.target.value)
    }
    const changeNeverExpire = () => {
        setNeverExpire(!neverExpire)
    }

    const [state, setState] = useState("")
    const [msg, setMsg] = useState("")
    const submit = async () => {
        setState("loading")
        setMsg("")
        if (!(originalURL.startsWith("http://") || originalURL.startsWith("https://"))) {
            setState("warning")
            setMsg("Original URL must start with \"http://\" or \"https://\"")
            return
        }
        if (originalURL.length < 8 || originalURL.length > 8182) {
            setState("warning")
            setMsg("Original URL isn't a valid URL")
            return
        }

        let data = {
            original_url: originalURL,
            duration_hours: neverExpire ? 0 : expireHours,
        }

        try {
            const responseData = await CreateURL(data)
            const expireAt = new Date(responseData.expired_at).toLocaleString()
            setState("success")
            setMsg(`${window.location.origin}/${responseData.short_path}`)
        } catch (e) {
            if (typeof e === "number") {
                switch (e) {
                    case 400:
                        setState("warning")
                        setMsg("Unsupport data, you may do something wrong")
                        break
                    case 500:
                        setState("error")
                        setMsg("Server internel error, please try again later or contact website admin")
                        break
                    default:
                        setState("error")
                        setMsg("Unknow error, please contact website admin")
                }
            } else {
                setState("error")
                setMsg(String(e))
            }
        }
    }

    return (
        <Box sx={{
            display: "flex",
            flexDirection: "column",
            gap: 5,

            width: "300px"
        }}
        >
            <Typography variant="h5">
                Generate Short URL:
            </Typography>

            <TextField
                label="Original URL"
                placeholder="https://example.com/xxxx"
                value={originalURL}
                onChange={changeOriginalURL}
            >
            </TextField>

            <FormControl fullWidth>
                <InputLabel id="expire">Expire</InputLabel>
                <Select
                    labelId="expire"
                    label="Expire"
                    value={expireHours}
                    onChange={changeExpireHours}
                    disabled={neverExpire}
                >
                    <MenuItem value={12}>12 Hours</MenuItem>
                    <MenuItem value={24}>24 Hours</MenuItem>
                    <MenuItem value={168}>7 Days</MenuItem>
                    <MenuItem value={720}>1 Month</MenuItem>
                    <MenuItem value={4320}>6 Months</MenuItem>
                    <MenuItem value={8760}>1 Year</MenuItem>
                    <MenuItem value={43800}>5 Years</MenuItem>
                </Select>
            </FormControl>

            <FormControlLabel control={
                <Switch checked={neverExpire} onChange={changeNeverExpire} />
            } label="Never Expire" />

            <Box>
                <Button
                    variant="contained"
                    onClick={submit}
                    loading={state === "loading"}
                >
                    Submit
                </Button>
            </Box>

            {state === "success" && (
                <Alert severity="success">
                    {msg}
                </Alert>
            )}
            {state === "warning" && (
                <Alert severity="warning">
                    Warning: {msg}
                </Alert>
            )}
            {state === "error" && (
                <Alert severity="error">
                    Error: {msg}
                </Alert>
            )}
        </Box>
    )
}

export default URL