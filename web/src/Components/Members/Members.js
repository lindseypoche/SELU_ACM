import React, { useState } from 'react';
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import classnames from 'classnames';
import './Members.css';
import MemberTable from '../Table/Table.js';

const Members = (props) => {
  const [activeTab, setActiveTab] = useState('1');

  const toggle = tab => {
    if(activeTab !== tab) setActiveTab(tab);
  }

  return (
    <div className="containerMember">
      <h1>Members</h1>
      <Nav className="tabamc" tabs>
        <NavItem className="navTabs">
          <NavLink
            id="memberTabs"
            className={classnames({ active: activeTab === '1' })}
            onClick={() => { toggle('1'); }}
          >
            Active Members
          </NavLink>
        </NavItem>
        <NavItem className="navTabs">
          <NavLink
            id="memberTabs"
            className={classnames({ active: activeTab === '2' })}
            onClick={() => { toggle('2'); }}
          >
            Inactive Members
          </NavLink>
        </NavItem>
      </Nav>
      <TabContent activeTab={activeTab}>
        <TabPane className="body" tabId="1">
          <MemberTable/>
        </TabPane>
        <TabPane className="body" tabId="2">
          <MemberTable/>
        </TabPane>
      </TabContent>
    </div>
  );
}

export default Members;