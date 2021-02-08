import React, { Component } from 'react';
import { Button, Form, FormGroup, Label, Input, FormText } from 'reactstrap';
import './Home.css';
import HomeCaro from './HomeCaro.js';
import ParallaxContainer from './ParallaxContainer';
import discord from '../image/Discord.png';


class Home extends Component {

  render() {
    return (
      // <div className="container"> 
      //   <div className="caro"><HomeCaro/></div>
      // </div>
      
      <div><div className="discord-thing"><a href="https://discord.gg/g6bQXFMjs3"><img src={discord}/></a></div>      <ParallaxContainer/></div>


    )
  }
}

export default Home;