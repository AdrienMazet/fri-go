import { Arwes, createTheme, ThemeProvider } from "arwes"
import React from "react"

type Props = {
  children: JSX.Element | JSX.Element[]
}

const Layout = ({ children }: Props) => (
  <ThemeProvider theme={createTheme({})}>
    <Arwes
      animate
      background="/images/background.jpg"
      pattern="/images/glow.png"
    >
      <main>{children}</main>{" "}
    </Arwes>
  </ThemeProvider>
)

export default Layout
