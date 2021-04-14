import React from 'react';
import axios from 'axios';
import { Parallax } from 'react-parallax';
import Events from "../../Components/Events/Events.js"
// import Featured from ".../Components/Featured/Featured.js"

export class EventsPage2 extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      events: [],
      eventsIsLoaded: false,
      eventsError: null,
      featured: {}, 
      featuredIsLoaded: false,
      featuredError: null,
      message: {},
      attachment: {}
    };
  }

  componentDidMount() {

    const acmFeaturedChannel = "817106404842143805" // events channel
    // const cakeFeaturedChannel = "814350227544604692"

    // fetch Featured data 
    axios.get("http://localhost:8080/featured/" + acmFeaturedChannel).then((response) => {
      this.setState({ 
        featured: response.data, 
        message: response.data.message,
        attachment: response.data.message.attachments,
        featuredIsLoaded: true,
       });
      console.log("featured: ", this.state.featured);
      console.log("message: ", this.state.message);
      console.log("attachment: ", this.state.attachment);
    }, 
      (featuredError) => {
        this.setState({
          featuredIsLoaded: true, 
          featuredError
        });
      }
    );

    // Fetch events data
    axios.get('http://localhost:8080/events').then((response) => {
      this.setState({ 
        events: response.data, 
        eventsIsLoaded: true,
       });
      console.log("events: ", this.state.events);
    }, 
      (eventsError) => {
        this.setState({
          eventsIsLoaded: true, 
          eventsError
        });
      }
    );

  }

  render() {
    const { eventsError, eventsIsLoaded, events, featuredError, featuredIsLoaded, featured, message, attachment } = this.state; 

      // return both featured and events 
        return (
            <>
            { !featuredError ? (
                <>
                  {/* <p>{message.id} : {attachment.url}</p> */}
                  {/* <Featured key={featured.channel_id} message={message} attachment={attachment} /> */}

                  <Parallax key={message.channel_id} blur={3} bgImage={attachment.url} bgImageAlt="never trust a bunny" strength={200} >
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
                        <p>
                            {message.content}
                        </p>
                    </div>
                  </div>
              </Parallax> 


                </>
              ) : (
                <div>No featured event found</div>
              )
            }

            { !eventsError ? (
                <Events events={events} />
            ) : (
                <div>No events found</div>
            )
            }
            </>
        )
  }
}

export default EventsPage2;