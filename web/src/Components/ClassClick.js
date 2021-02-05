import React, { Component } from 'react'

// EVENT HANDLING IN A CLASS COMPONENTS
class ClassClick extends Component {

    clickHandler(){
        console.log('Clicked the button')
    }

    render() {
        return (
            <div>
                <button onClick = {this.clickHandler}> Click me</button>
            </div>
        )
    }
}

export default ClassClick
