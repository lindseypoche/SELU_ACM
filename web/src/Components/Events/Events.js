import './Events.css'
import { Link } from "react-router-dom";
import { toDateFormat, isExpiring, getRemainingTime } from "../../Utils/timing.js"
import { useState,useEffect} from 'react';
import axios from 'axios';

const readingTime = require('reading-time');

const Events = () => {

  const [events, setEvent] = useState({});
  const [eventIsLoaded, setEventIsLoaded] = useState(false)
  const [avatar, setAvatar] = useState([]); 
  const [author, setAuthor] = useState([]);

  useEffect(() => {
      getEvent();
  }, [])

  const getEvent = () => {
     axios.get(`http://localhost:8081/api/events`)
         .then(((response) => {  
             setEvent(response.data);
             setEventIsLoaded(true);
             setAvatar(response.data.author.avatar);
             setAuthor(response.data.author);
        }))
        .catch(error => console.log(error))
    }

    if (!eventIsLoaded) {
      return <div className="App">Loading...</div>;
    }

    return (
        <>
        {
            (events.length > 0 ? (
              <div className="container">
                <div className="upcoming__wrapper">
                  <span>Upcoming Events</span>
                </div>
              {events.map((event) => (
                    <div className="column">
                      <article className="article">
                        {
                            isExpiring(event.timestamp) ? (
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
                            <img className="attachment__image" src={event.attachments.url} />
                          )
                        }
                        <span className="starting__date">Event {getRemainingTime(event.timestamp)}</span>
                        <h2 className="article__title">{event.content.substring(0, 40)}</h2>
                        <p className="article__excerpt">
                          {/* <ReactMarkdown> */}
                            {event.content.substring(0, 100)}
                          {/* </ReactMarkdown> */}
                        </p>
                        </Link>
                        <div className="user__info__wrapper">
                          <div id="avatar">
                            <img className="avatar__image" src={avatar.image_url} />
                          </div>
                          <div className="info">
                            <p className="username">{author.username}</p>
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