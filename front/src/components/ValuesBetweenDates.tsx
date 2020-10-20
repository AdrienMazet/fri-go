import { Heading, Table } from "arwes"
import React, { useState } from "react"
import dates from "../assets/dates.json"
import SelectOption from "../components/Select"

const sensorTypes = [
  { name: "Temperature", prop: "Temperature", unit: "°C" },
  { name: "Pression", prop: "Pressure", unit: "Pa" },
  { name: "Vents solaires", prop: "Wind", unit: "Parsec/s" },
]

type Props = {
  sensorValues: any
}

const ValuesBetweenDates = ({ sensorValues }: Props) => {
  const [startDate, setStartDate] = useState(dates[0])
  const [endDate, setEndDate] = useState(dates[0])

  const getDatesBetweenDates = (startDate, endDate) => {
    let retDates = []
    let sDate = new Date(startDate)
    let eDate = new Date(endDate)
    if (sDate > eDate) {
      sDate = new Date(endDate)
      eDate = new Date(startDate)
    }
    while (sDate < eDate) {
      const currentDate = new Date(sDate).toISOString().substring(0, 10)
      if (dates.find(date => date === currentDate))
        retDates = [...retDates, currentDate]
      sDate.setDate(sDate.getDate() + 1)
    }
    if (new Date(startDate) > new Date(endDate)) {
      retDates = [...retDates, startDate]
    } else {
      retDates = [...retDates, endDate]
    }

    return retDates
  }

  const getMaxValuesLength = date => {
    return Math.max(
      sensorValues[sensorTypes[0].prop.toLowerCase()][date].length,
      sensorValues[sensorTypes[1].prop.toLowerCase()][date].length,
      sensorValues[sensorTypes[2].prop.toLowerCase()][date].length
    )
  }

  return (
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
      {getDatesBetweenDates(startDate, endDate).map(date => (
        <div key={date}>
          <Heading node="h3">{date}</Heading>
          <Table
            animate
            headers={sensorTypes.map(
              sensorType => sensorType.name + " (" + sensorType.unit + ")"
            )}
            dataset={Array.from(
              { length: getMaxValuesLength(date) },
              (_, i) => i
            ).map(index => [
              sensorValues[sensorTypes[0].prop.toLowerCase()][date][
                index
              ].toFixed(2),
              sensorValues[sensorTypes[1].prop.toLowerCase()][date][
                index
              ].toFixed(2),
              sensorValues[sensorTypes[2].prop.toLowerCase()][date][
                index
              ].toFixed(2),
            ])}
          />
        </div>
      ))}
    </div>
  )
}

export default ValuesBetweenDates
