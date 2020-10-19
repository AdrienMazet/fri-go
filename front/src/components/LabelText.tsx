import React from "react"
import { withStyles } from "arwes"
import Words from "arwes/lib/Words"
import myTheme from "../theme/base_theme"
const styles = theme => ({
  labelText: {
    width: 140,
    display: "inline-flex",
    paddingRight: "1rem",
    textAlign: "right",
    fontFamily: myTheme.font.fontFamily.regular,
    fontSize: 30,
    height: myTheme.font.baseSize + 20,
    [`@media (max-width: ${theme.responsive.small + 1}px)`]: {
      display: "block",
      textAlign: "left",
      paddingBottom: ".5rem",
      width: "100%",
    },
  },
})

const LabelText = withStyles(styles)(({ classes, label, required }) => {
  return (
    <Words layer="primary">
      <span className={classes.labelText}>
        {label}
        {required ? (
          <span>
            <Words layer="alert">*</Words>
          </span>
        ) : (
          ""
        )}
      </span>
    </Words>
  )
})

export default LabelText
