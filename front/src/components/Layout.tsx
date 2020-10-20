import { Arwes, createTheme, Image, ThemeProvider, Header, Button } from "arwes"
import { Link } from "gatsby"
import React, { useEffect, useState } from "react"
import Context from "../context/Context"
import variantThemes from "../theme/variant_themes"
import ThemeSelector from "./ThemeSelector"

type Props = {
  children: JSX.Element | JSX.Element[]
  name?: string
  header?: boolean
}

const Layout = ({ children, name, header }: Props) => {
  const [themeName, setThemeName] = useState(typeof window !== 'undefined' && localStorage.getItem("themeName"))
  const [theme, setTheme] = useState(createTheme())

  const updateTheme = (themeName?) => {
    let myTheme
    if (themeName === "luke" || !themeName) {
      myTheme = createTheme()
      setTheme(myTheme)
      themeName = "luke"
      setThemeName(themeName)
      localStorage.setItem("themeName", themeName)
    } else {
      myTheme = createTheme(variantThemes[themeName])
      setTheme(myTheme)
      setThemeName(themeName)
      localStorage.setItem("themeName", themeName)
    }
  }

  useEffect(() => {
    if (!themeName) {
      updateTheme()
    } else {
      updateTheme(themeName)
    }
  }, [themeName])

  return (
    <Context.Provider value={{ theme, themeName, updateTheme }}>
      <ThemeProvider theme={theme}>
        <Arwes
          animate
          background="/images/background.jpg"
          pattern="/images/glow.png"
        >
          {header && (
            <div style={{ padding: 20 }}>
              <Header animate>
                <div
                  style={{
                    display: "flex",
                    justifyContent: "space-between",
                    alignItems: "center",
                  }}
                >
                  <div>
                    <h1 style={{ margin: 0 }}>{`Spatioport${
                      name ? " : " + name : "s"
                    }`}</h1>
                    {name && (
                      <Link to="/">
                        <Button
                          animate
                          style={{ marginTop: 15, marginBottom: 15 }}
                        >
                          Accueil
                        </Button>
                      </Link>
                    )}
                  </div>
                  <div style={{ display: "flex" }}>
                    <ThemeSelector
                      style={{ marginRight: 15, marginTop: name ? 12 : 0 }}
                    />
                    {name && (
                      <Image
                        animate
                        resources={`/images/spatioports/${name}.jpg`}
                        style={{ width: 100, marginRight: 15 }}
                      />
                    )}
                  </div>
                </div>
              </Header>
            </div>
          )}
          <main>{children}</main>
        </Arwes>
      </ThemeProvider>
    </Context.Provider>
  )
}

export default Layout
