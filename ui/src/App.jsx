import {
	createTheme,
	ThemeProvider,
} from "@mui/material"
import Bar from './componets/Bar'
import Info from "./componets/Info"
import { useState } from "react"


function App() {
	const [mode, setMode] = useState("dark")
	const theme = createTheme({
		palette: {
			mode: mode,
		}
	})
	const changeMode = () => {
		setMode(mode === "dark" ? "light" : "dark")
	}

	return (
		<ThemeProvider theme={theme}>
			<Bar mode={mode} changeMode={changeMode} />
			<Info />
		</ThemeProvider>
	)
}

export default App