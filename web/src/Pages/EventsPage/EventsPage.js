import { Component } from 'react';
import {Link} from "react-router-dom"
import axios from 'axios';
import './EventsPage.css'
import Events from "../../Components/Events/Events.js";

export class EventsPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      events: [],
      eventsIsLoaded: false,
      eventsError: null,
      pin: {
        timestamp: '',
        edited_timestamp: '',
        start_time: '',
        title: '',
        content: '',
        attachments: [],
        author: {
          avatar: {}
        },
      },
      pinIsLoaded: false,
      pinError: null
    };
  }

  componentDidMount() {
    const eventsChannel = "817106404842143805" 

    // fetch latest pin 
    axios.get("http://localhost:8081/api/pins/channel?id="+eventsChannel).then((response) => {
      this.setState({ 
        pin: response.data, 
        pinIsLoaded: true,
       });
    }, 
      (pinError) => {
        this.setState({
          pinError: true
        });
      }
    );

  }

  render() {
    const { eventsError, eventsIsLoaded, events, pinError, pinIsLoaded, pin } = this.state; 

    if (!pinIsLoaded) {
      return <div className="App">Loading...</div>;
    }

    return (

    <div className="main">
      {/* <div className="banner">@kevin joined the discord server! Woot!</div> */}
      { pin.id != "" ? (
          <div className="featured__event">

            <div className="ribbon pinned__ribbon">
              <p>📍</p>
            </div>

            <img className="featured__image" src={pin.attachments[0].url} /> 
            <div className="avatar__title__wrapper">
              <h2 className="featured__title">
                <img src={pin.author.avatar.image_url}></img>
                {pin.title} 
                <Link to={`/event/${pin.id}`}>
                  <span> learn more</span>
                </Link>
              </h2>
            </div>
          </div>
        ) : ('')
      }

      <Events />

    </div>
    )
  }
}

export default EventsPage;