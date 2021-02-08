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
                jhkasklcvnrfklvndfv <br/>
                useDebugValue(feverv <br/>
                wesfv <br/>
                sdf <br/>
                vs <br/>
                fdv <br/>
                sdfv <br/>
                sdfv <br/>
                sd <br/>
                fv <br/>
                sdfv) <br/>
                jhkasklcvnrfklvndfv <br/>
                useDebugValue(feverv <br/>
                wesfv <br/>
                sdf <br/>
                vs <br/>
                fdv <br/>
                sdfv <br/>
                sdfv <br/>
                sd <br/>
                fv <br/>
                sdfv) <br/>
                jhkasklcvnrfklvndfv <br/>
                useDebugValue(feverv <br/>
                wesfv <br/>
                sdf <br/>
                vs <br/>
                fdv <br/>
                sdfv <br/>
                sdfv <br/>
                sd <br/>
                fv <br/>
                sdfv) <br/>
                jhkasklcvnrfklvndfv <br/>
                useDebugValue(feverv <br/>
                wesfv <br/>
                sdf <br/>
                vs <br/>
                fdv <br/>
                sdfv <br/>
                sdfv <br/>
                sd <br/>
                fv <br/>
                sdfv) <br/>
                jhkasklcvnrfklvndfv <br/>
                useDebugValue(feverv <br/>
                wesfv <br/>
                sdf <br/>
                vs <br/>
                fdv <br/>
                sdfv <br/>
                sdfv <br/>
                sd <br/>
                fv <br/>
                sdfv) <br/>
            </div>
        )
    }
}

export default ParallaxContainer
