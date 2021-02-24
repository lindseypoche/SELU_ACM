import React, { Component } from "react";
import { Button, Form, FormGroup, Label, Input, FormText } from "reactstrap";
import "./Footer.css";

class Footer extends Footer {
    render() {
        return(
            
            <div classname= "footie">
                <div class = "feetpics">
                    <img src="/image/SoutheasternLogo.png" alt="SouthTransparent"></img>
                </div>
                <div classname = "address">
                    <li>Southeastern Louisiana University</li>
                    <li>Hammond, Louisiana 70402</li>
                    <li></li>
                    <li></li>
                    <a href="mailto:lindsey.poche-2@southeastern.edu">Questions or Comments</a>
                </div>
                
            </div>

        );
    }
}
export default Footer;