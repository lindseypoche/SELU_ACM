import React from 'react'

// EVENT HANDLING
// react events are CamelCase rather than lowercase
function FunctionClick() {
    function clickHandler(){
        console.log('Button Clicked')
    }
    return (
        <div>
            <button onClick = {clickHandler}> Click</button>
        </div>
    )
}

export default FunctionClick