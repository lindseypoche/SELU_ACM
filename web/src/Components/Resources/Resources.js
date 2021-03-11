import React, { Component } from 'react';
import { Button, Form, FormGroup, Label, Input, FormText } from 'reactstrap';
import './Resources.css';
import Fay from '../image/Schedules/FAY125Schedule.png'
import CSTB from '../image/Schedules/CSTB2026Schedule.png'



class Resources extends Component {

  render() {
    return (
      <div className="resourcesPage"> 
        <h1>Resources</h1>
        <center>
        <br></br>
        <p>Tutoring for CMPS 161, 280, 290, and 390 is available in the Computer Science major labs located in CSTB 2026 & FAYARD Hall 125. Both labs are open Monday through Thursday 9:30 AM to 9:00 PM.</p>
        <br></br>
        <h3>CSTB 2026 Schedule</h3>
        <div className = "image">
        <img src={CSTB} alt ="CSTB Schedule"  ></img>
        </div>
        <br></br>
        
        <h3>Fayard 125 Schedule</h3>
        <div className = "image">
        <img src={Fay} alt ="Fay Schedule"></img>
        </div>
        </center>
        <br></br>
      </div>
    )
  }
}

export default Resources;