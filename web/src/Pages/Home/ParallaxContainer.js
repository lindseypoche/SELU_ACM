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
                <div className="deta">
                <h1>About Us</h1>
                <h5>The Association for Computing Machinery (ACM) at Southeastern Louisiana University is a student chapter of the ACM Organization. We are a professional and social group of Computer Science and IT majors. Join and get to know fellow students, exchange tips, and learn from peers and professionals. Expand your network within the Computer Science community.</h5>
                <br></br>
                <br></br>
                <h1>National Info</h1>
                <h5 >"ACM brings together computing educators, researchers, and professionals to inspire dialogue, share resources, and address the field's challenges. As the world’s largest computing society, ACM strengthens the profession's collective voice through strong leadership, promotion of the highest standards, and recognition of technical excellence. ACM supports the professional growth of its members by providing opportunities for life‐long learning, career development, and professional networking."
                    <br></br>
                -The ACM Organization</h5>
                <br></br>
                <br></br>
                <h1>Southeastern's ACM</h1>
                <h5>ACM brings together students in the Computer Science department, giving them the opportunities they need to connect them with industry leaders, secure internships, meet new friends, and stay successful in their journey through  coding development. We host bi-weekly meetings which include pizza, resume workshops, code workshops, and much more. At our Distinguished Lecturer series, we invite guests from local Businesses and professional groups that work with Southeastern Alumni to start their careers and offer internships for those not ready to graduate. We also host an Annual Crawfish boil in the Spring semester!</h5>
            
                <br></br>
                <br></br>
                <h1>Distingushed Lecturers</h1>
                <h5><br></br>Calvin Farve, CEO and Owner of Envoc;
                <br></br>Willam Assaf, SQL Guru at Microsoft;
                <br></br>Kevin Cefalu, Netchex Powershell Developer;
                <br></br>Chris, MMR;
                <br></br>Marcel, IT Director at MMR.</h5>
                
            </div>
            </div>
        )
    }
}

export default ParallaxContainer
