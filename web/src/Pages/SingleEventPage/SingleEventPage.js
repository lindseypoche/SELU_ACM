import React, { useState, useEffect } from 'react';
import axios from 'axios';
import ReactMarkdown from 'react-markdown'
import './SingleEventPage.css';
import { Button, Avatar, Image, Box } from 'grommet';
import { FaRegComment, FaRegHeart, FaGrinTongueWink } from 'react-icons/fa';
import {toDateFormat, isExpiring, getRemainingTime} from "../../Utils/timing.js"

const readingTime = require('reading-time');

// class SingleEventPage extends Component {
const SingleEventPage = ({ match }) => {

    const [event, setEvent] = useState({});
    const [eventIsLoaded, setEventIsLoaded] = useState(false)

    useEffect(() => {
        getEvent();
    }, [])

    const getEvent = () => {
       axios.get(`http://localhost:8081/api/events/${match.params.id}`)
           .then(((response) => {  
               setEvent(response.data);
               setEventIsLoaded(true);
          }))
          .catch(error => console.log(error))
      }

      if (!eventIsLoaded) {
        return <div className="App">Loading...</div>;
      }

        return (
            <div className="container">
                <div className="single-container">
                    <div className="title-container">
                        <div className="title">
                            <ReactMarkdown>
                                {event.content.substring(0, 30)}
                            </ReactMarkdown>
                        </div>
                    </div>
                    <div className="details-container">
                        <Avatar 
                            src={event.author.avatar.image_url}
                            size="medium"
                            background="lightgray"
                            margin="xxsmall"
                        >
                        </Avatar>
                        <Button
                            background="none"
                            color="green"
                            font-size="small"
                            margin="small"
                        >{event.author.username}</Button>
                        <p className="text-issue">{toDateFormat(event.timestamp)} • {readingTime(event.content).text}</p>
                    </div>
                    <div className="img-container">
                        <Box
                            width="xxlarge"
                        >
                            { event.attachments != null ? (
                            <Image
                                fit="contain"
                                src={event.attachments.url}
                            />
                            ) : (
                                <div>
                                    <p>No Image Available</p>
                                </div>
                            )
                            }
                        </Box>
                    </div>
                    <div className="content-container">
                        <div className="body-container">
                            <ReactMarkdown>
                                {event.content}
                            </ReactMarkdown>
                        </div>
                    </div>
                </div>
            </div>
        );
}
export default SingleEventPage;