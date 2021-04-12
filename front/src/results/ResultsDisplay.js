import React from 'react';
import TableItem from './TableItem';
import './ResultsDisplay.css';

const ResultsDisplay = (props) => {
    const {list} = props

    return (
        <div id="results-display">
            <table>
                <thead>
                    <tr>
                        <th>Company</th>
                        <th>{list && list[0].measureTag}</th>
                        <th>Ending Period</th>
                    </tr>
                </thead>
                <tbody>
                    {list && list.map((item, idx) => {
                        return(
                            <TableItem item={item} idx={idx} />
                        )
                    })}
                </tbody>
            </table>
        </div>
    );
}

export default ResultsDisplay;