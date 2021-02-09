import React, { Component } from 'react';
import { Button, Form, FormGroup, Label, Input, FormText } from 'reactstrap';
import './About.css';



class About extends Component {

  render() {
    return (
      <div className="aboutPage"> 
      
        <div className=""><h1>About</h1></div>
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

export default About;