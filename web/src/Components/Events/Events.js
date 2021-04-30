import React, { useState, useEffect } from 'react';
import axios from 'axios'
import {Link} from "react-router-dom"
import { toDateFormat, isExpiring, getRemainingTime } from "../../Utils/timing.js"
import './Events.css'
import Author from './Author.js'

const readingTime = require('reading-time');

const Events = () => {

  const [events, setEvents] = useState([])
  const [eventsIsLoaded, setEventsIsLoaded] = useState(false)
  const [eventsErr, setEventsErr] = useState(false)

  useEffect(() => {
      getEvents();
  }, [])

  const getEvents = () => {
     axios.get(`http://localhost:8081/api/events`)
         .then(((response) => {  
             setEvents(response.data);
             setEventsIsLoaded(true);
        }))
        .catch(error => {
          setEventsErr(true)
        })
    }

    if (!eventsIsLoaded) {
      return <div className="App">Loading...</div>;
    }

    if (eventsErr) {
      return <div>error getting events</div>
    }

    console.log("events: ", events)

    return (
        <>
        {
            (events.length > 0 ? (
              <div className="container">
                <div className="upcoming__wrapper">
                  <span>Upcoming Events</span>
                </div>
              {events.map((event) => (
                    <div className="column" key={event.id}>
                      <article className="article">
                        {
                            isExpiring(event.start_time) ? (
                            <div className="ribbon expiration__ribbon">
                              <p>🕚</p>
                            </div>
                          ) : ('')
                        }
                        <Link to={`/event/${event.id}`} style={{color: "inherit", textDecoration: "none"}} >
                        {
                          event.attachments == null ? (
                           <div className="no__photo">
                             <h2>No Photo Available</h2>
                             <p></p>
                           </div> 
                          ) : (

                           <img className="attachment__image" src={event.attachments[0].url} />

                          )
                        }
                        <span className="starting__date">Event {getRemainingTime(event.start_time)}</span>
                        <h2 className="article__title">{event.title}</h2>
                        <p className="article__excerpt">
                            {event.content.substring(0, 170)}
                        </p>
                        </Link>
                        <div className="user__info__wrapper">
                          <div id="avatar">
                            <Author author={event.author} />
                            {/* <img className="avatar__image" src={event.author.avatar.image_url} /> */}
                          </div>
                          <div className="info">
                            {/* <p className="username">{event.author.username}</p> */}
                            <p className="date"> {toDateFormat(event.timestamp)} • {readingTime(event.content).text}</p>
                          </div>
                        </div>
                      </article>
                    </div>
              ))}
             </div>
               ) : (
                //  make template event?
                'No Events To Show'
               ) 
            )
        }
        </>
    )
}

export default Events;