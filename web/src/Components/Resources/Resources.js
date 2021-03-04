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