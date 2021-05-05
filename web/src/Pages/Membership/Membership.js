import React, { Component } from 'react';
import './Membership.css';
import square from "../../Media/Images/Icons/squareUp.png";
import form from "../../Media/Forms/SELU ACM Membership Registration Form.pdf";
import pdf from "../../Media/Images/Icons/pdf.png";
import acm from '../../Media/Images/Icons/association-for-computing-machinery-logo.png';


class Membership extends Component {

  render() {
    return (
    <div className="membershipPage">
        <div className= "membershipPara">
          <h1>Membership</h1>
          <br></br>
          <br></br>
          <div className = "acm">
            <img src = {acm}></img>
          </div>
          <br></br>
          <br></br>
          <h3>Want to Join the ACM?</h3>
          <br></br>
          <p1>Download these forms and either bring them to a local ACM Officer or email them to us at:</p1>
          <a href="mailto:acm@selu.edu"> acm@selu.edu</a>
          <br></br>
          <br></br>
          <br></br>
          <br></br>
          <div className = "forms">
            <a href = {form}>
              <img src={pdf} alt = "pdf"></img>
            </a>
          </div>
          <div className = "form1">
            <a src = {pdf}>
            <p>SELU ACM Membership Form</p>
            </a>
          </div>
          <br></br>
          <br></br>
          <p>Then pay membership fees via Square or in person with a local SELU ACM Officer.</p>
          <br></br>
          <br></br>
          <div className="square">
            <a href="https://checkout.square.site/merchant/1VRKAABSW92ZE/checkout/I2XGX7WFRSNVVTXCC4HTHDXE">
              <img src={square} alt="SquareImg"  ></img>
            </a>
          </div>
          <br></br>
   
        </div>
      </div>
    )
  }
}

export default Membership;