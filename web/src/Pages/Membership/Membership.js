import React, { Component } from 'react';
import './Membership.css';
import square from "../../Media/Images/squareUp.png";
import form from "../../Media/Forms/SELU ACM Membership Registration Form.pdf";

class Membership extends Component {

  render() {
    return (

        <div className="membershipPage">

            <div className= "membershipPara">

                <h1>Membership</h1>
                <br></br>
                <br></br>
                <div className="form">
                  <p>import forms and then have acm email to send forms to</p>
                  <br></br>
                </div>
                <div className="square">
                  <a href="https://checkout.square.site/merchant/1VRKAABSW92ZE/checkout/I2XGX7WFRSNVVTXCC4HTHDXE">
                  <center><img src={square} alt="SquareImg"  ></img></center>
                  </a>
                </div>

                <br></br>
                <p>add paypal??</p>
                
            </div>

        </div>
    )
  }
}

export default Membership;