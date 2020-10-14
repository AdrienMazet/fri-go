import { Button, Header, Project, Words } from "arwes"
import React from "react"
import Layout from "../components/Layout"
import conf from "../../../config/conf.json"
import SpatioportCard from "../components/SpatioportCard"

const IndexPage = () => (
  <Layout>
    <div style={{ padding: 20 }}>
      <Header animate>
        <h1 style={{ margin: 0 }}>Spatioports</h1>
      </Header>
    </div>
    <h3 style={{ marginLeft: 20, marginBottom: 100 }}>
      <Words animate>A proximité de votre planète :</Words>
    </h3>
    <div style={{display: "flex", alignItems: "center", justifyContent:"center"}}>
      {conf.Airports.map(airport => (
        <SpatioportCard name={airport} />
      ))}
    </div>
  </Layout>
)

export default IndexPage
