import React from 'react'
import {DateParser} from '../utils/DateParser'

const TableItem = (props) => {
    const {item, idx} = props
    const {companyName, value, periodEndDate} = item

    const endPeriod = DateParser(periodEndDate)
    return (
        <>
            <tr key={idx}>
                <td>{companyName}</td>
                <td>{value}</td>
                <td>{endPeriod}</td>
            </tr>
        </>
    )
}

export default TableItem
