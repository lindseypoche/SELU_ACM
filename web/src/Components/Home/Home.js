import React, { Component } from 'react';
import { Button, Form, FormGroup, Label, Input, FormText } from 'reactstrap';
import './Home.css';
import HomeCaro from './HomeCaro.js';


class Home extends Component {

  render() {
    return (
      <div className="container"> 
        <div className="caro"><HomeCaro/></div>
      </div>
    )
  }
}

export default Home;