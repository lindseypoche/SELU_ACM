import React, { Component } from "react";
import { Button, Form, FormGroup, Label, Input, FormText } from "reactstrap";
import "./Footer.css";

class Footer extends Footer {
    render() {
        return(
            
            <div classname= "footie" role="contentinfo" aria-label="Main Footer">
                <div class = "feetpics">
                    <a href="/">
                        <img src="/image/SoutheasternLogo.png" alt="SELUTransparent"></img>
                    </a>
                </div>
                <div classname = "address">
                    <li>Southeastern Louisiana University</li>
                    <li>Hammond, Louisiana 70402</li>
                    <li></li>
                    <li></li>
                    <a href="mailto:acm@selu.edu">Questions or Comments</a>
                    <li>University Policies</li>
                    <li>Accessiblilty Information</li>
                    <li>Chapter Policies and Bylaws</li>
                </div>
                {/*add
                http://www.southeastern.edu/
                https://www.southeastern.edu/resources/policies/index.html
                https://www.southeastern.edu/resources/accessibility/index.html
                https://www.acm.org/chapters/about-chapters
                https://www.acm.org/chapters/chapter-policies
                */}
            </div>

        );
    }
}
export default Footer;