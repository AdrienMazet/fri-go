import { Header, Words, Button } from "arwes"
import React from "react"
import Layout from "../components/Layout"
import { Link } from "gatsby"

const NotFoundPage = () => (
  <Layout>
    <Header animate>
      <h1 style={{ margin: 0 }}>404: Not Found</h1>
    </Header>
    <div
      style={{ display: "flex", flexDirection: "row", alignItems: "center" }}
    >
      <p>
        <Words animate layer="alert">
          You went throught the wrong portal, go back to your root galaxy quick.
        </Words>
      </p>
      <Link to="/">
        <Button animate style={{ marginLeft: 15 }}>
          Take me home
        </Button>
      </Link>
    </div>
  </Layout>
)

export default NotFoundPage
