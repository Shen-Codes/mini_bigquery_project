import React from 'react'

const AcctToSearch = (props) => {

    const handleSelect = e => {
        props.setState(prev=> ({
            ...prev,
            acct: e.target.value
        }))
    }

    return (
        <div>
            <input list="account-list" type="text">
                <datalist id="account-list">
                    <option value="Revenue" onSelect={handleSelect}/>
                    <option value="Cost of Goods Sold" onSelect={handleSelect}/>
                    <option value="Net Income or Loss" onSelect={handleSelect}/>
                </datalist>
            </input>
        </div>
    )
}

export default AcctToSearch