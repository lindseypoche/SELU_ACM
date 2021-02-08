import React, { Component } from 'react';
import { Button, Form, FormGroup, Label, Input, FormText } from 'reactstrap';
import './Home.css';
import HomeCaro from './HomeCaro.js';
import ParallaxContainer from './ParallaxContainer';
import discord from '../image/Discord.png';


class Home extends Component {

  render() {
    return (
<<<<<<< HEAD
      // <div className="container"> 
      //   <div className="caro"><HomeCaro/></div>
      // </div>
      
      <div><div className="discord-thing"><a href="https://discord.gg/g6bQXFMjs3"><img src={discord}/></a></div>      <ParallaxContainer/></div>

=======
      <div className="container"> 
        <h1>Home</h1>
        <div className="caro"><HomeCaro/></div>
      </div>
>>>>>>> 65042d32914d03d2b1b7c98be36403d335bd8a49
    )
  }
}

export default Home;