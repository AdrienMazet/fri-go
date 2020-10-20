import { Heading, Words } from "arwes"
import React, { useState } from "react"
import dates from "../assets/dates.json"
import SelectOption from "../components/Select"

const sensorTypes = [
  { name: "Temperature", prop: "Temperature", unit: "°C" },
  { name: "Pression", prop: "Pressure", unit: "Pa" },
  { name: "Vents solaires", prop: "Wind", unit: "Parsec/s" },
]

type Props = {
  averageSensorValues: any
}

const AverageValues = ({ averageSensorValues }: Props) => {
  const [dateAverage, setDateAverage] = useState(dates[0])
  return (
    <div style={{ width: "30%", marginLeft: 25 }}>
      <Heading node="h1">Valeurs moyennes</Heading>
      <SelectOption
        label="Date : "
        name="date"
        onChange={e => setDateAverage(e.target.value)}
        options={dates}
      />
      {sensorTypes.map(sensorType => (
        <div
          key={sensorType.prop}
          style={{ display: "flex", alignItems: "center" }}
        >
          <h4>
            <Words animate>{sensorType.name + " : "}</Words>
          </h4>
          <p style={{ marginLeft: 15 }}>
            <Words animate>
              {averageSensorValues[dateAverage][sensorType.prop]
                .toFixed(2)
                .toString() +
                " " +
                sensorType.unit}
            </Words>
          </p>
        </div>
      ))}
    </div>
  )
}

export default AverageValues
