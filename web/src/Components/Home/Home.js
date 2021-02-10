import React, { Component } from 'react';
import { Button, Form, FormGroup, Label, Input, FormText } from 'reactstrap';
import './Home.css';
import HomeCaro from './HomeCaro.js';
import ParallaxContainer from './ParallaxContainer';
import discord from '../image/Discord.png';



class Home extends Component {

  render() {
    return (

      
      <div>
<<<<<<< HEAD
        <div className="discord-thing"><a href="https://discord.gg/g6bQXFMjs3"><img src={discord}/></a>
        </div>      
      <ParallaxContainer/>
      </div>
=======
        <div className="discord-thing"><a href="https://discord.gg/g6bQXFMjs3"><img src={discord}/></a></div>      
        <ParallaxContainer/>
       </div>
>>>>>>> Dev


    )
  }
}

export default Home;