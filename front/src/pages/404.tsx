import { Header, Words, Button } from "arwes"
import React from "react"
import Layout from "../components/Layout"
import { Link } from "gatsby"

const NotFoundPage = () => (
  <Layout header={false}>
    <Header animate style={{ padding: 20 }}>
      <h1 style={{ margin: 0 }}>404: Page non trouvée</h1>
    </Header>
    <div
      style={{
        display: "flex",
        flexDirection: "row",
        alignItems: "center",
        marginLeft: 15,
      }}
    >
      <p>
        <Words animate layer="alert">
          Vous avez emprunté le mauvais portail, retournez vite à votre galaxie.
        </Words>
      </p>
      <Link to="/">
        <Button animate style={{ marginLeft: 15 }}>
          Ramenez moi à la maison
        </Button>
      </Link>
    </div>
  </Layout>
)

export default NotFoundPage
