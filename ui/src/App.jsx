import {
	createTheme,
	ThemeProvider,
} from "@mui/material"
import Bar from './componets/Bar'
import Info from "./componets/Info"

const theme = createTheme({
	palette: {
		mode: "light",
	}
})

function App() {
	return (
		<ThemeProvider theme={theme}>
			<Bar />
			<Info />
		</ThemeProvider>
	)
}

export default App