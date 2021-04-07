import React from 'react'

const Year = (props) => {

    const handleChange = e => {
        props.setState(prev => ({
            ...prev,
            year: e.target.value 
        }))
    }

    return (
        <div>
            <label htmlFor="years" >Year</label>
            <input list="year-list" type="number" id="years" onChange={handleChange}/>
                <datalist id="year-list">
                    <option value={2010}/>
                    <option value={2011}/>
                    <option value={2012}/>
                    <option value={2013}/>
                    <option value={2014}/>
                    <option value={2015}/>
                    <option value={2016}/>
                    <option value={2017}/>
                </datalist>          
        </div>
    )
}

export default Year