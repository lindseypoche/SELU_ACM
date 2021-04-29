import React, { Component } from "react";
import "./Home.css";

import ParallaxContainer from "./ParallaxContainer";
import VideoPlayback from '../../Media/Videos/VideoPlayback.js';


class Home extends Component {
  render() {
    return (
      
      <div className="homePage">

        <VideoPlayback />
        <div className="discord-thing">
          
          
        </div>
        <div className="paraHome"><ParallaxContainer /></div>
        
        <div className="takeUpSpace"></div>
      </div>
    );
  }
}

export default Home;
