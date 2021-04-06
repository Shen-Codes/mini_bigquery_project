import React from 'react'
import AcctToSearch from './AcctToSearch'
import AmountSlider from './AmountSlider'
import Year from './Year'

const SearchBar = (props) => {
    return (
        <div>
            <AcctToSearch setState={props.setState}/>
            <AmountSlider min={props.searchParams.min} max={props.searchParams.max} setState={props.setState} />
            <Year setState={props.setState}/>
        </div>
    )
}

export default SearchBar
