import {
  Arwes,
  Button,
  createTheme,
  Header,
  Project,
  ThemeProvider,
  Words,
} from "arwes"
import React from "react"
// import { Link } from "gatsby"

// import Layout from "../components/layout"
// import Image from "../components/image"
// import SEO from "../components/seo"

// const IndexPage = () => (
//   <Layout>
//     <SEO title="Home" />
//     <h1>Hi people</h1>
//     <p>Welcome to your new Gatsby site.</p>
//     <p>Now go build something great.</p>
//     <div style={{ maxWidth: `300px`, marginBottom: `1.45rem` }}>
//       <Image />
//     </div>
//     <Link to="/page-2/">Go to page 2</Link> <br />
//     <Link to="/using-typescript/">Go to "Using TypeScript"</Link>
//   </Layout>
// )

const IndexPage = () => (
  <ThemeProvider theme={createTheme({})}>
    <Arwes animate background="/img/background.jpg" pattern="/img/glow.png">
      <div style={{ padding: 20 }}>
        <Header animate>
          <h1 style={{ margin: 0 }}>Arwes - Cyberpunk UI Framework</h1>
        </Header>
      </div>
      <h3>
        <Words animate>A cyberpunk UI project</Words>
      </h3>
      <div style={{ padding: "20px" }}>
        <Button animate disabled>
          Futuristic
        </Button>{" "}
        <Button animate>Cyberpunk</Button>{" "}
        <Button animate layer="success">
          <i className="mdi mdi-chemical-weapon" /> Sci Fi
        </Button>{" "}
        <Button animate layer="alert">
          High Tech <i className="mdi mdi-robot" />
        </Button>
      </div>
      <p>
        <Words animate>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusamus,
          amet cupiditate laboriosam sunt libero aliquam, consequatur alias
          ducimus adipisci nesciunt odit? Odio tenetur et itaque suscipit atque
          officiis debitis qui. Lorem ipsum dolor sit amet, consectetur
          adipisicing elit. Accusamus, amet cupiditate laboriosam sunt libero
          aliquam, consequatur alias ducimus adipisci nesciunt odit? Odio
          tenetur et itaque suscipit atque officiis debitis qui. Lorem ipsum
          dolor sit amet, consectetur adipisicing elit. Accusamus, amet
          cupiditate laboriosam sunt libero aliquam, consequatur alias ducimus
          adipisci nesciunt odit? Odio tenetur et itaque suscipit atque officiis
          debitis qui.
        </Words>
      </p>
      <p>
        <Words animate layer="success">
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusamus,
          amet cupiditate laboriosam sunt libero aliquam, consequatur alias
          ducimus adipisci nesciunt odit? Odio tenetur et itaque suscipit atque
          officiis debitis qui.
        </Words>
      </p>
      <p>
        <Words animate layer="alert">
          With animations based on SciFi and designs from high technology
        </Words>
      </p>
      <div style={{ padding: 20 }}>
        <Project animate header="PROJECT, OFFICIA DESERUNT ANIM ID EST LABORUM">
          {anim => (
            <p>
              <Words animate show={anim.entered}>
                Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do
                eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut
                enim ad minim veniam, quis laboris nisi ut aliquip ex. Duis aute
                irure. Consectetur adipisicing elit, sed do eiusmod tempor
                incididunt ut labore et dolore magna aliqua. Ut enim ad minim
                veniam, quis nostrud.
              </Words>
            </p>
          )}
        </Project>
      </div>
    </Arwes>
  </ThemeProvider>
)

export default IndexPage
