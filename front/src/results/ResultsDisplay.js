import React from 'react';
import TableItem from './TableItem';

const ResultsDisplay = (props) => {
    const {list} = props

    return (
        <div>
            <table>
                <thead>
                    <tr>
                        <th>Company</th>
                        <th>{list[0].measureTag}</th>
                        <th>Ending Period</th>
                    </tr>
                </thead>
                <tbody>
                    {list && list.map((item, idx) => {
                        <TableItem item={item} idx={idx} />
                    })}
                </tbody>
            </table>
        </div>
    );
}

export default ResultsDisplay;