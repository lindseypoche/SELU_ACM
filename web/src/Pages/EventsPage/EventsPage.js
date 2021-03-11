import React, { useState, useEffect } from "react";
import axios from "axios"
import { Parallax } from 'react-parallax';
import Featured from "../../Components/Featured/Featured.js"
import Events from "../../Components/Events/Events.js"

const EventsPage = () => {

        let eventsRoute = "http://localhost:8080/events";
        let featuredEventRoute = "http://localhost:8080/featured/817106404842143805";
        // const eventsReq = axios.get(events);
        // const featuredReq = axios.get(featuredEvent);
      
      const [featuredHook, setFeaturedHook] = useState([]);
      const [eventsHook, setEventsHook] = useState([]);
      useEffect(() => {
        axios.all(
                [axios.get(featuredEventRoute),
                axios.get(eventsRoute)])
            .then(axios.spread((featuredResponse, eventsResponse) => {  
                 setFeaturedHook(featuredResponse.data);
                 setEventsHook(eventsResponse.data);
                console.log("featured: ", featuredResponse.data);
                console.log("events: ", eventsResponse.data);
            }))
            .catch(error => console.log(error));
      }, [])

        return (
            <>
              <Parallax key={featuredHook.channel_id} blur={3} bgImage={featuredHook.message.attachments.url} bgImageAlt="never trust a bunny" strength={200}>
                  <div style={
                      {
                          height: '700px',
                          position: 'relative',
                          margin: '0 auto',
                      }
                      }>
                    <div style={
                        {
                            height: 'auto', 
                            margin: '0 200px',
                            padding: '10px',
                            position: 'absolute',
                            top: '65%', 
                            transform: 'translateY(-35%)',
                            backgroundColor:'#fff', 
                        }
                    }>
                        <p style={
                            {
                            }
                        }>
                            {featuredHook.message.content}
                        </p>
                    </div>
                  </div>
              </Parallax>
                {/* <Featured key={featuredHook.channel_id} featured={featuredHook} /> */}

              <Events events={eventsHook} />
            </>
        )
}

export default EventsPage;