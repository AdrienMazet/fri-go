import { Button, Image, Project } from "arwes"
import { Link } from "gatsby"
import React from "react"

type Props = {
  name: string
}

const SpatioportCard = ({ name }: Props) => (
  <div style={{ padding: 20, width: "23%" }}>
    <Project animate header={`Spatioport : ${name}`}>
      <div>
        <Image animate resources={`/images/spatioports/${name}.jpg`} />
      </div>
      <div style={{ padding: "20px" }}>
        <Link to={`/spatioport/${name}`}>
          <Button
            animate
            style={{
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
            }}
          >
            <p style={{ fontSize: 30 }}>Relevés spaciaux</p>
          </Button>
        </Link>
      </div>
    </Project>
  </div>
)

export default SpatioportCard
