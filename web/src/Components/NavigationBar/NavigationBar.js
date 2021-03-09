import React, { useState } from 'react';
import acmlogo from '../image/association-for-computing-machinery-logo.png'
import {
  Collapse,
  Navbar,
  NavbarToggler,
  NavbarBrand,
  Nav,
  NavItem,
  NavLink,
  UncontrolledDropdown,
  DropdownToggle,
  DropdownMenu,
  DropdownItem,
  NavbarText
} from 'reactstrap';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
import './NavigationBar.css';

const NavigationBar = (props) => {
  const [isOpen, setIsOpen] = useState(false);

  const toggle = () => setIsOpen(!isOpen);

  return (
    <div className="navacmbg">
      <Navbar className="navacm" expand="md">
        <NavbarBrand>
          <NavLink className="acm" tag={Link} to='/'><img style = {{height: '50px', width: '50px', marginRight: '10px'}} src = {acmlogo} />Association for Computing Machinery</NavLink>
        </NavbarBrand>
        <NavbarToggler onClick={toggle} />
        <Collapse isOpen={isOpen} navbar>
          <Nav className="mr-auto" navbar>
          <NavItem>
              <NavLink className="acm" tag={Link} to="/officers">Officers</NavLink>
            </NavItem>
            <NavItem>
              <NavLink className="acm" tag={Link} to="/events">Events</NavLink>
            </NavItem>
            <NavItem>
              <NavLink className="acm" tag={Link} to="/resources">Resources</NavLink>
            </NavItem>
            <NavItem>
              <NavLink className="acm" tag={Link} to="/calendar">Calendar</NavLink>
            </NavItem>
          </Nav>
          {/* <Nav className="rightSide" navbar>
          <UncontrolledDropdown nav inNavbar>
              <DropdownToggle className="acm" nav caret>
                Admin
              </DropdownToggle>
              <DropdownMenu right>
                <DropdownItem className="acm" tag={Link} to="/members">
                  Members
                </DropdownItem>
                
                <DropdownItem className="acm" tab={Link} to="/events">
                  Events
                </DropdownItem>
                <DropdownItem className="acm">
                  Images
                </DropdownItem>
              </DropdownMenu>
              </UncontrolledDropdown>
              <NavItem>
                <NavLink className="acm" tag={Link} to="/join">Join</NavLink>
              </NavItem>
              <NavItem>
                <NavLink className="acm" tag={Link} to="/login">Login</NavLink>
              </NavItem>
            </Nav> */}
        </Collapse>
      </Navbar>
    </div>
  );
}

export default NavigationBar;