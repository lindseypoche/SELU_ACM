import React, { Component } from 'react';
import { Button, Form, FormGroup, Label, Input, FormText } from 'reactstrap';
import './Login.css';


class Login extends Component {

  render() {
    return (
      <div className="LogInPage">
          <Form className="LogInForm">
            <h2 className="logheader">Loginn</h2>
            <br />
            <FormGroup>
              <Label for="exampleEmail">Username</Label>
              <Input type="email" name="email" id="exampleEmail"/>
            </FormGroup>
            <FormGroup>
              <Label for="examplePassword">Password</Label>
              <Input type="password" name="password" id="examplePassword"/>
            </FormGroup>
            <Button className="logButton" size="lg" block>Login</Button>
          </Form>
        </div>
    )
  }
}

export default Login;