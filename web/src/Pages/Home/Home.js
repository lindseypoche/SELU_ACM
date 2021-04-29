import React, { Component } from "react";
import "./Home.css";

import ParallaxContainer from "./ParallaxContainer";
import discord from "../../Media/Images/Discord.png";
import VideoPlayback from '../../Media/Videos/VideoPlayback.js';


class Home extends Component {
  render() {
    return (
      
      <div className="homePage">

        <VideoPlayback />
        <div className="discord-thing">
          <a href="https://discord.gg/g6bQXFMjs3">
            <img src={discord} />
          </a>
        </div>
        <ParallaxContainer />
        <div className="takeUpSpace"></div>
      </div>
    );
  }
}

export default Home;
