import { Words, Heading } from "arwes"
import React, { useContext, useState } from "react"
import dates from "../assets/dates.json"
import Layout from "../components/Layout"
import SelectOption from "../components/Select"
import Context from "../context/Context"

const sensorTypes = [
  { name: "Temperature", prop: "Temperature", unit: "°C" },
  { name: "Pression", prop: "Pressure", unit: "Pa" },
  { name: "Vents solaires", prop: "Wind", unit: "Parsec/s" },
]

type Props = {
  pageContext: {
    data: {
      name: string
      sensorValues: any
      averageSensorValues: any
    }
  }
}

const SpatioportTemplate = (props: Props) => {
  const x = useContext(Context)
  const [startDate, setStartDate] = useState(dates[0])
  const [endDate, setEndDate] = useState(dates[0])
  const [dateAverage, setDateAverage] = useState(dates[0])
  const { name, sensorValues, averageSensorValues } = props.pageContext.data
  console.log(x)

  console.log(sensorValues)
  return (
    <Layout name={name} header>
      <div style={{ display: "flex", alignItems: "flex-start" }}>
        <div
          style={{ width: "30%", marginLeft: 25, borderRight: `solid black` }}
        >
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
        <div style={{ width: "60%", marginLeft: 150 }}>
          <Heading node="h1">Valeurs par périodes</Heading>
          <div style={{ display: "flex", alignItems: "center" }}>
            <SelectOption
              label="Du : "
              name="period"
              onChange={e => setStartDate(e.target.value)}
              options={dates}
            />
            <div style={{ marginLeft: 50 }}>
              <SelectOption
                label="Au : "
                name="period"
                onChange={e => setEndDate(e.target.value)}
                options={dates}
              />
            </div>
          </div>
        </div>
      </div>
    </Layout>
  )
}

export default SpatioportTemplate
