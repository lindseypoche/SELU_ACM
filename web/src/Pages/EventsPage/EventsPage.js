import React, { FC, useState, useEffect } from 'react'
import axios from 'axios';
import './EventsPage.css'
import Events from "../../Components/Events/Events.js"


export class EventsPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      events: [],
      eventsIsLoaded: false,
      eventsError: null,
      pin: {
        timestamp: '',
        edited_timestamp: '',
        content: '',
        attachments: {},
        author: {
          avatar: {}
        },
      },
      pinIsLoaded: false,
      pinError: null
    };
  }

  componentDidMount() {
    // const acmFeaturedChannel = "817106404842143805" // events channel
    // const acmFeaturedChannel = "814350227544604692" // cakes events channel

    // fetch latest pin 
    axios.get("http://localhost:8081/api/pins/channel?id=817106404842143805").then((response) => {
      this.setState({ 
        pin: response.data, 
        // attachments: response.data.attachments,
        // author: response.data.author,
        // avatar: response.data.author.avatar,
        // content: response.data.content,
        pinIsLoaded: true,
       });
      // console.log("pin: ", this.state.pin);
      // console.log("attachment: ", this.state.attachment);
    }, 
      (pinError) => {
        this.setState({
          pinError: true
        });
      }
    );

    // featch all active events
    axios.get('http://localhost:8081/api/events').then((response) => {
      this.setState({ 
        events: response.data, 
        eventsIsLoaded: true,
       });
      console.log("events: ", this.state.events);
    }, 
      (eventsError) => {
        this.setState({
          eventsError: true
        });
      }
    );

  }

  render() {
    const { eventsError, eventsIsLoaded, events, pinError, pinIsLoaded, pin } = this.state; 

    if (!eventsIsLoaded || !pinIsLoaded) {
      return <div className="App">Loading...</div>;
    }

    // console.log("events have loaded: ", events)

    return (

    <div>
      {/* <div className="banner">@kevin joined the discord server! Woot!</div> */}
      { pin.id != "" ? (
          <div className="featured__event">

            <div className="ribbon pinned__ribbon">
              <p>📍</p>
            </div>

            <img className="featured__image" src={pin.attachments.url} /> 
            <div className="avatar__title__wrapper">
              <h2 className="featured__title"><img src={pin.author.avatar.image_url}></img>{pin.content.substring(0, 100)} <a href="#">(learn more)</a></h2>
            </div>
          </div>
        ) : ('')
      }
    {
      (!eventsError ? (
        <Events events={events} />
      ): (
        <div>No upcomming events</div>
      ))
    }       

    </div>
    )
  }
}

export default EventsPage;