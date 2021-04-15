import React, { Component } from 'react';
import './Resources.css';
import Fay from '../../Media/Images/Schedules/FAY125Schedule.png'
import CSTB from '../../Media/Images/Schedules/CSTB2026Schedule.png'

class Resources extends Component {

  render() {
    return (
      <div className="resourcesPage">
        <div className="resourcesPara">
          <h1>Resources</h1>
          <center>
            <br></br>
            <p>Tutoring for CMPS 161, 280, 290, and 390 is available in the Computer Science major labs located in CSTB 2026 & FAYARD Hall 125. Both labs are open Monday through Thursday 9:30 AM to 9:00 PM.</p>
            <br></br>
            <h3>CSTB 2026 Schedule</h3>
            <div className="image">
              <img src={CSTB} alt="CSTB Schedule"  ></img>
            </div>
            <br></br>

            <h3>Fayard 125 Schedule</h3>
            <div className="image">
              <img src={Fay} alt="Fay Schedule"></img>
            </div>
          </center>
          <br></br>
        </div>
      </div>
    )
  }
}

export default Resources;