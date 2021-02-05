import React, { Component } from 'react'

class Counter extends Component {

    // The only place you can assign this.state is in the constructor.
    // Anywhere else must use the setState method.

    constructor(props) {
        super(props)
    
        this.state = {
             count: 0
        }
    }
    
    increment(){
        // this.state.count = this.state.count + 1
        // this.setState(
        //     {
        //         count: this.state.count + 1
        //     },
        //     () => {
        //         console.log('Callback value - ', this.state.count)
        //     } // anytime you need to execute some code after changing the state,
        //     //do it here in the second parameter (callback method) of the setState function
        // )

        this.setState((prevState,props) => ({
            count: prevState.count + 1
        }))
        console.log(this.state.count)
    }

    incrementFive(){
        this.increment()
        this.increment()
        this.increment()
        this.increment()
        this.increment()
    }

    render() {
        return (
            <div>
               <div> Count - {this.state.count} </div>
               <button onClick= {() => this.incrementFive()}> Increment</button>
            </div>
        )
    }
}

export default Counter
