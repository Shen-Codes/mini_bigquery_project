import React from 'react'

const AcctToSearch = (props) => {
    return (
        <div>
            <input list="account-list" type="text">
                <datalist id="account-list">
                    <option value="Revenue" onSelect={console.log("data")}/>
                    <option value="Cost of Goods Sold" onSelect={console.log("data")}/>
                    <option value="Net Income or Loss" onSelect={console.log("data")}/>
                </datalist>
            </input>
        </div>
    )
}

export default AcctToSearch