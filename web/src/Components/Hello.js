import React from 'react'

const Hello = () => {
    // JSX VERSION
    return (
        <div className = 'dummyClass'>
            <h1> Hello Person</h1>
        </div>
    )

    // // WITHOUT JSX
    // return React.createElement('div', 
    //     {id: 'hello', className: 'dummyClass'}, 
    //     React.createElement('h1', null, 'Hello Vishwas')
    // )

}

export default Hello