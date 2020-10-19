import { Header, Image } from "arwes"
import React, { useState } from "react"
import dates from "../assets/dates.json"
import Layout from "../components/Layout"

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

  // display image
  console.log(sensorValues)
  console.log(averageSensorValues)
  return (
    <Layout name={name}>
      <div style={{ display: "flex", alignItems: "flex-start" }}>
        <div style={{ width: "40%", marginLeft: 25 }}></div>
        <div style={{ width: "60%" }}>
          <select multiple onChange={event => console.log(event)}>
            {dates.map(date => (
              <option value={date}>{date}</option>
            ))}
          </select>
        </div>
      </div>
    </Layout>
  )
}

export default SpatioportTemplate

// () => {
//   const [startDate, setStartDate] = useState(new Date());
//   const [endDate, setEndDate] = useState(null);
//   const onChange = dates => {
//     const [start, end] = dates;
//     setStartDate(start);
//     setEndDate(end);
//   };
//   return (
//     <DatePicker
//       selected={startDate}
//       onChange={onChange}
//       startDate={startDate}
//       endDate={endDate}
//       selectsRange
//       inline
//     />
//   );
// };
