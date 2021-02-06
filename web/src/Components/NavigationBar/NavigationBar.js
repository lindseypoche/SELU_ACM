import React, { useState } from 'react';
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
      <Navbar className="navacm" light expand="md">
        <NavbarBrand>
          <NavLink className="acm" tag={Link} to='/Home'>ACM</NavLink>
        </NavbarBrand>
        <NavbarToggler onClick={toggle} />
        <Collapse isOpen={isOpen} navbar>
          <Nav className="mr-auto" navbar>
            <NavItem>
              <NavLink >Events</NavLink>
            </NavItem>
            <NavItem>
              <NavLink>About</NavLink>
            </NavItem>
          </Nav>
          <Nav className="rightSide" navbar>
              <NavItem>
                <NavLink href="/components/">Join</NavLink>
              </NavItem>
              <NavItem>
                <NavLink tag={Link} to="/Login">Login</NavLink>
              </NavItem>
            </Nav>
        </Collapse>
      </Navbar>
    </div>
  );
}

export default NavigationBar;