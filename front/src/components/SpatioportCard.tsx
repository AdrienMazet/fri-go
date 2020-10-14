import { Button, Project, Words } from "arwes"
import React from "react"

type Props = {
  name: string
}

const SpatioportCard = ({ name }: Props) => (
  <div style={{ padding: 20, width: "23%" }}>
    <Project animate header={`Spatioport : ${name}`}>
      {anim => (
        <p>
          <Words animate show={anim.entered}>
            Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do
            eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim
            ad minim veniam, quis laboris nisi ut aliquip ex. Duis aute irure.
            Consectetur adipisicing elit, sed do eiusmod tempor incididunt ut
            labore et dolore magna aliqua. Ut enim ad minim veniam, quis
            nostrud.
          </Words>
        </p>
      )}
      <p>
        <Words animate>Moyenne de la journée </Words>
      </p>
      <p>
        <Words animate>Temperature : 50 °C </Words>
      </p>
      <p>
        <Words animate layer="success">
          Pression : 2 Pascals
        </Words>
      </p>
      <p>
        <Words animate layer="alert">
          Vents solaires : 75km/h
        </Words>
      </p>
      <div style={{ padding: "20px" }}>
        <Button animate>Plus de détails</Button>
      </div>
    </Project>
  </div>
)

export default SpatioportCard
