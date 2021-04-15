import React, { Component } from 'react'
import './Parallax.css';

{/*
import Home from './Components/Home/Home.js';
import Calendar from './Components/Calendar/Calendar.js';
import Officers from './Components/Officers/Officers.js'
import Events from './Components/Events/Events.js';
import Resources from './Components/Resources/Resources';
*/}

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
                <h1>About Us</h1>
                <p1>The Association for Computing Machinery (ACM) at Southeastern Louisiana University is a student chapter of the ACM Organization. We are a professional and social group of Computer Science and IT majors. Join and get to know fellow students, exchange tips, and learn from peers and professionals. Expand your network within the Computer Science community.</p1>
                <br></br>
                <br></br>
                <h1>National Info</h1>
                <p1>"ACM brings together computing educators, researchers, and professionals to inspire dialogue, share resources, and address the field's challenges. As the world’s largest computing society, ACM strengthens the profession's collective voice through strong leadership, promotion of the highest standards, and recognition of technical excellence. ACM supports the professional growth of its members by providing opportunities for life‐long learning, career development, and professional networking."
                    <br></br>
                    -The ACM Organization</p1>
            </div>
        )
    }
}

export default ParallaxContainer
