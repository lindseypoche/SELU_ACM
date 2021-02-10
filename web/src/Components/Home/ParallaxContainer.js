import React, { Component } from 'react'
import './Parallax.css';

class ParallaxContainer extends Component { 
    constructor(props){
        super(props);
        this.myRef = React.createRef();
    }

    componentDidMount() {
        window.addEventListener("scroll", (ev) => {
            if (this.myRef.current!=null){
                this.myRef.current.style.transform = 'translateY(' + -document.documentElement.scrollTop + 'px)';
            }
        })
    }

    render() {
        return (
            <div className="parallax" ref={this.myRef}>

            </div>
        )
    }
}

export default ParallaxContainer
