import React, { Component } from "react";
import { Button, Form, FormGroup, Label, Input, FormText } from "reactstrap";
import "./Home.css";
import HomeCaro from "./HomeCaro.js";
import ParallaxContainer from "./ParallaxContainer";
import discord from "../image/Discord.png";
import VideoPlayback from '../video/VideoPlayback.js'


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
        
      </div>
    );
  }
}

export default Home;
