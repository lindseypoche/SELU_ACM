import React, { useState } from 'react';
import acmlogo from '../../Media/Images/association-for-computing-machinery-logo.png';
import { Nav, Navbar} from "react-bootstrap";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  useHistory
} from "react-router-dom";
import './NavigationBar.css';

const NavigationBar = (props) => {
  const [isOpen, setIsOpen] = useState(false);
  let history = useHistory();

  const toggle = () => setIsOpen(!isOpen);

  return (
    <div className="navacmbg">
      < Navbar className="navacm" collapseOnSelect expand="md" bg="dark" fixed="top" variant="dark" >
        <Navbar.Brand className="acm" onClick={()=> history.push("/home")}><img style={{ height: '50px', width: '50px', marginRight: '10px' }} src={acmlogo} />Association for Computing Machinery</Navbar.Brand>
        <Navbar.Toggle aria-controls="responsive-navbar-nav" />
        <Navbar.Collapse id="responsive-navbar-nav">
          <Nav className="mr-auto">
            <Nav.Link className="acm" onClick={()=> history.push("/officers")}>Officers</Nav.Link>
            <Nav.Link className="acm" onClick={()=> history.push("/events")}>Events</Nav.Link>
            <Nav.Link className="acm" onClick={()=> history.push("/resources")}>Resources</Nav.Link>
            <Nav.Link className="acm" onClick={()=> history.push("/calendar")}>Calendar</Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Navbar >
    </div>
  );
}

export default NavigationBar;