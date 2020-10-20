import { Frame, withStyles } from "arwes"
import React, { useContext } from "react"
import context from "../context/Context"
import myTheme from "../theme/base_theme"

const styles = theme => ({
  wrapper: {
    display: "inline-block",
    marginTop: 5,
  },
  select: {
    display: "inline-block",
    margin: [5, 0],
    height: 30,
    width: 80,
    border: "none",
    outline: "none",
    boxShadow: "none",
    cursor: "pointer",

    backgroundColor: myTheme.color.background.main,
    color: theme.color.primary.base,
    fontSize: 16,
    lineHeight: 30,
  },
  frame: {
    width: "auto",
    padding: "0 3px",
    display: "inline-block",
  },
})

const ThemeSelector = withStyles(styles)(({ classes, ...props }) => {
  const { updateTheme, themeName } = useContext(context)

  const setTheme = e => {
    updateTheme(e.target.value)
  }
  return (
    <span {...props}>
      <div className={classes.wrapper}>
        <Frame
          show={true}
          animate={true}
          corners={10}
          className={classes.frame}
        >
          <select onChange={setTheme} className={classes.select}>
            <option value="luke" selected={themeName === "luke"}>
              Luke
            </option>
            <option value="windu" selected={themeName === "windu"}>
              Windu
            </option>
            <option value="vader" selected={themeName === "vader"}>
              Vader
            </option>
            <option value="yoda" selected={themeName === "yoda"}>
              Yoda
            </option>
            <option value="rey" selected={themeName === "rey"}>
              Rey
            </option>
          </select>
        </Frame>
      </div>
    </span>
  )
})

export default ThemeSelector
