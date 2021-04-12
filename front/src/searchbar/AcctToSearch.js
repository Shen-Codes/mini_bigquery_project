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
                    <option value="Revenues" />
                    <option value="CostofGoodsSold" />
                    <option value="NetIncomeorLoss" />
                </datalist>      
        </div>
    )
}

export default AcctToSearch