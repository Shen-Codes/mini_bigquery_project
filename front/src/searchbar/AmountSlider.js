import React from 'react'

const AmountSlider = (props) => {

    const handleChange = e => {
        e.preventDefault();
        props.setState(prev => ({
            ...prev,
            [e.target.name]: e.target.value
        }))
    }

    return (
        <div>
            <label htmlFor="min">Min amount</label>
            <input 
                type="number" 
                id="min" 
                name="min"
                value={props.min} 
                min="-9999999999" 
                max="99999999999"
                onChange={handleChange}
            />
            <label htmlFor="max">Max amount</label>
            <input 
                type="number" 
                id="max" 
                name="max"
                value={props.max} 
                min="9999999999" 
                max="99999999999"
                onChange={handleChange}
            />
        </div>
    )
}

export default AmountSlider