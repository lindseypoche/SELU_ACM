import React from 'react';
import axios from 'axios';
import { Parallax } from 'react-parallax';
import Events from "../../Components/Events/Events.js"
import Featured from "../../Components/Featured/Featured.js"

export class EventsPage2 extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      events: [],
      eventsIsLoaded: false,
      featured: {}, 
      featuredIsLoaded: false,
    };
  }

  componentDidMount() {

    // fetch Featured data 
    axios.get('http://localhost:8080/featured/817106404842143805').then((response) => {
      this.setState({ 
        featured: response.data, 
        featuredIsLoaded: true,
       });
      console.log("featured: ", this.state.featured);
    }, 
      (error) => {
        this.setState({
          featuredIsLoaded: true, 
          error
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
      (error) => {
        this.setState({
          eventsIsLoaded: true, 
          error
        });
      }
    );

  }


  render() {
    const { error, eventsIsLoaded, events, featuredIsLoaded, featured } = this.state; 

    if (error) {
      return <div>Error: {error.message}</div>;
    } else if (!eventsIsLoaded || !featuredIsLoaded) {
      return <div>Loading...</div>
    } else {
        return (
            <>
            {/* <Featured featured={featured} /> */}

          <Parallax key={this.state.featured.channel_id} blur={3} bgImage={this.state.featured.message.attachments.url} bgImageAlt="never trust a bunny" strength={200} >
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
                            {this.state.featured.message.content.substring(0, 250)}
                        </p>
                    </div>
                  </div>
              </Parallax> 



            <Events events={events} />
            </>
        )
    }
  }
}

export default EventsPage2;