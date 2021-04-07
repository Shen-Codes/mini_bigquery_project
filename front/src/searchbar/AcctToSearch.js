import React from 'react'

const AcctToSearch = (props) => {

    const handleSelect = e => {
        props.setState(prev=> ({
            ...prev,
            account: e.target.value
        }))
    }

    return (
        <div>
            <label htmlFor="accounts">Account to Search</label>
            <input list="account-list" type="text" id="accounts" onChange={handleSelect}/>
                <datalist id="account-list">
                    <option value="Revenue" />
                    <option value="Cost of Goods Sold" />
                    <option value="Net Income or Loss" />
                </datalist>      
        </div>
    )
}

export default AcctToSearch