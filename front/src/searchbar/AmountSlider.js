import React from 'react'

const AmountSlider = (props) => {
    return (
        <div>
            <input type="number" value={props.min} min="" max=""/>
            <input type="number" value={props.max} min="" max=""/>
        </div>
    )
}

export default AmountSlider