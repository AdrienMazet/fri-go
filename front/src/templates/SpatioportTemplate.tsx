import React from "react"
import AverageValues from "../components/AverageValues"
import Layout from "../components/Layout"
import ValuesBetweenDates from "../components/ValuesBetweenDates"

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
  const { name, sensorValues, averageSensorValues } = props.pageContext.data

  return (
    <Layout name={name} header>
      <div style={{ display: "flex", alignItems: "flex-start" }}>
        <AverageValues averageSensorValues={averageSensorValues} />
        <ValuesBetweenDates sensorValues={sensorValues} />
      </div>
    </Layout>
  )
}

export default SpatioportTemplate
