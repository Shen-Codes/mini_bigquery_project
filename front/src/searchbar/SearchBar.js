import React from 'react'
import AcctToSearch from './AcctToSearch'
import AmountSlider from './AmountSlider'
import Year from './Year'
import "./SearchBar.css"

const SearchBar = (props) => {
    return (
        <div id="search-bar">
            <AcctToSearch setState={props.setState}/>
            <AmountSlider min={props.searchParams.min} max={props.searchParams.max} setState={props.setState} />
            <Year setState={props.setState}/>
            <button onClick={props.fetch}>Search</button>
        </div>
    )
}

export default SearchBar
