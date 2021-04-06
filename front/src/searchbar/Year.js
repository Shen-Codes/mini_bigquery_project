import React from 'react'

const Year = (props) => {


    return (
        <div>
            <input list="year-list" type="string"/>
                <datalist id="year-list">
                    <option value="2010" onSelect={() => props.setState(prev => ({...prev, year: 2010}))}/>
                    <option value="2011" onSelect={() => props.setState(prev => ({...prev, year: 2011}))}/>
                    <option value="2012" onSelect={() => props.setState(prev => ({...prev, year: 2012}))}/>
                    <option value="2013" onSelect={() => props.setState(prev => ({...prev, year: 2013}))}/>
                    <option value="2014" onSelect={() => props.setState(prev => ({...prev, year: 2014}))}/>
                    <option value="2015" onSelect={() => props.setState(prev => ({...prev, year: 2015}))}/>
                    <option value="2016" onSelect={() => props.setState(prev => ({...prev, year: 2016}))}/>
                    <option value="2017" onSelect={() => props.setState(prev => ({...prev, year: 2017}))}/>
                </datalist>          
        </div>
    )
}

export default Year