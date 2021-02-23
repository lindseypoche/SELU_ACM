import React, { Component } from 'react';
import { Button, Form, FormGroup, Label, Input, FormText } from 'reactstrap';
import './Events.css';
import EventCard from '../Card/EventCard.js';


class Events extends Component {

  render() {
    return (
      <div className="eventsPage"> 
        <h1>Events</h1>
        {/*
        <div className="eventContainer">
            <EventCard/>
            <EventCard/>
            <EventCard/>
        </div>


        
        */}
      </div>
    )
  }
}

export default Events;