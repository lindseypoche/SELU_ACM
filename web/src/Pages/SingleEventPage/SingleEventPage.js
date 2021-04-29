import React, { useState, useEffect } from 'react';
import axios from 'axios';
import ReactMarkdown from 'react-markdown'
import './SingleEventPage.css';
import { Button, Avatar, Image, Box, Carousel } from 'grommet';
import {toDateFormat, isExpiring, getRemainingTime} from "../../Utils/timing.js"
import Comments from '../../Components/Comments/Comments.js';

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
                            // width="xxlarge"
                            height="medium"
                            // width="xxxlarge"
                            overflow="hidden"
                        >

                        <Carousel fill="false" play="5000" controls="">
                        {/* <Carousel fill="true" play="5000" controls="false"> */}
                            { 
                            event.attachments.length > 0 ? (
                                event.attachments.map((attachment, i) => (
                                    // <Image key={i} fit="cover" src={attachment.url} />
                                    <Image key={i} src={attachment.url} />
                                ))
                            ) : (
                                <div>
                                    <p>No Image Available</p>
                                </div>
                            )
                            }
                            </Carousel>
                        </Box>
                    </div>

                    {
                        (event.message_reactions.count > 0 ?
                            (
                                <div className="emoji-container">
                                    <p className="reaction-title"></p>
                                    <div className="sticky">
                                    {
                                        (event.message_reactions.reactions.map((reaction => (
                                            <>
                                                <span> {reaction.emoji.name} </span>
                                            </>
                                        ))))
                                    }
                                    </div>
                                </div>
                            ) : (
                                ''
                            )
                        )
                    }

                    <div className="content-container">
                        <div className="body-container">
                            <ReactMarkdown>
                                {event.content}
                            </ReactMarkdown>
                        </div>
                    </div>

                    <div>
                        <Comments eventID={event.id}/>
                    </div>
                </div>

            </div>

        );
}
export default SingleEventPage;