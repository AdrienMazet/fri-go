const axios = require("axios")
const conf = require("../config/conf.json")
const dates = require("./src/assets/dates.json")
const sensorTypes = require("./src/assets/sensorTypes.json")

exports.createPages = async ({ actions, graphql, reporter }) => {
  const { createPage } = actions

  const spatioportTemplate = require.resolve(
    `./src/templates/SpatioportTemplate.tsx`
  )

  const url = "http://localhost:8080"

  const getAverageSensorsValues = (airport, date) =>
    axios.get(`${url}/${airport}/${date}/results/average`)

  const getSensorValuesBetweenDates = (airport, sensorType, date1, date2) =>
    axios.get(`${url}/${airport}/filter/${sensorType}/${date1}/${date2}`)

  conf.Airports.forEach(airport => {
    let sensorValues
    let averageSensorValues

    Promise.all(
      sensorTypes.map(async sensorType => {
        const response = await getSensorValuesBetweenDates(
          airport,
          sensorType,
          dates[0],
          dates[dates.length - 1]
        )
        sensorValues = { ...sensorValues, sensorType: response.data }
      })
    ).then(() => {
      Promise.all(
        dates.map(async date => {
          const res = await getAverageSensorsValues(airport, date)
          averageSensorValues = { ...averageSensorValues, date: res.data }
        })
      ).then(() =>
        createPage({
          path: `/spatioport/${airport}`,
          component: spatioportTemplate,
          context: {
            data: { name: airport, sensorValues, averageSensorValues },
          },
        })
      )
    })
  })
}
