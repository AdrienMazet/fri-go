import { Header } from "arwes"
import React from "react"
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
  const {name, sensorValues, averageSensorValues} = props.pageContext.data
  return (
    <Layout>
      <div style={{ padding: 20 }}>
        <Header animate>
          <h1 style={{ margin: 0 }}>{name}</h1>
        </Header>
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

// () => {
//   const [startDate, setStartDate] = useState(null);
//   return (
//     <DatePicker
//       selected={startDate}
//       onChange={date => setStartDate(date)}
//       includeDates={[new Date(), addDays(new Date(), 1)]}
//       placeholderText="This only includes today and tomorrow"
//     />
//   );
// };
