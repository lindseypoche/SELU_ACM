import React, { Component } from "react";
import { Button, Form, FormGroup, Label, Input, FormText } from "reactstrap";
import "./Footer.css";
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
  } from "react-router-dom";
import LOGO from '../image/SoutheasternLogo.png'
import ACM from '../image/association-for-computing-machinery-logo.png'

class Footer extends Component {
    render() {
        return(
            <div class= "footie">
                <br></br>
                <br></br>
                <div class = "Column1">
                    <div class = "feetpics">
                        <a href="https://www.southeastern.edu/">
                        <div class="selu" tag={Link} to='/'><img style = {{height: '55px', width: '50px', marginRight: '10px'}} src = {LOGO} /></div>
                        </a>
                        <br></br>
                       
                    <div class="acm" tag={Link} to='$'><img style = {{height: '55px', width: '50px', marginRight: '10px'}} src = {ACM} /></div>
                        
                    </div>
                        <div class = "address">

                            <br></br>
                            <br></br>
                            <br></br>
                            <div>Southeastern Louisiana University</div>
                            <div>Hammond, Louisiana 70402</div>
                            <div></div>
                        </div>
                    <br></br>
                </div>
                <div class = ' Column2'>
                <br></br>
                    <a href="mailto:acm@selu.edu">Questions or Comments</a>
                    <br></br>
                    <a href = "https://my-store-11562067.creator-spring.com/">Want Merch?</a>
                    <br></br>
                    <a href = "https://www.southeastern.edu/resources/policies/index.html">University Policies</a>
                    <br></br>
                    <a href = "https://www.southeastern.edu/resources/accessibility/index.html">Accessiblilty Information</a>
                    <br></br>
                    <a href = "https://www.acm.org/chapters/about-chapters"> Chapter Information</a>
                    <br></br>
                    <a href = "https://www.acm.org/chapters/chapter-policies">Chapter Policies and Bylaws</a>
                    <br></br>
                </div>

            </div>

        );
    }
}
export default Footer;