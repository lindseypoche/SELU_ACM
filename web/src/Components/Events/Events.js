import './Events.css'

const readingTime = require('reading-time');

function toDateFormat(unix) {
  var months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sept", "Oct", "Nov", "Dec"]
  var d = new Date(unix * 1000);
  var month = months[d.getMonth()];
  var f = month + " " + d.getDay().toString()
  return f
}

function isExpiring(exp_date) {
  var seconds_remaining = parseInt(exp_date) - Date.now()/ 1000
  if(seconds_remaining < 0) {
    return false
  }
  // if event is between 0 and 2 days 
  if(seconds_remaining < 86400*2) {
    return true
  }
  return false
}

function getRemainingTime(start_date) {
  var format = ""
  var delta = parseInt(start_date) - Date.now()/ 1000

  if(delta < 0) {
    return "ended"
  }
  if(delta < 3600) {
    return "starts in < 1 hour"
  }

  // calculate and subtract days
  var days = Math.floor(delta / 86400); 
  delta -= days * 86400;
  if(days < 2) {
    format = "starts in " + days + " day"
  } else {
    format = "starts in " + days + " days"
  }

  // calculate and subtract hours 
  var hours = Math.floor(delta / 3600) % 24;
  delta -= hours * 3600;
  if(hours < 2) {
    format += " and " + hours + " hour"
  } else {
    format += " and " + hours + " hours"
  }

  return format
}

const Events = ({ events }) => {

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
                  <p className="article__excerpt">{event.content.substring(0, 100)}</p>
                  <div className="user__info__wrapper">
                    <div id="avatar">
                      <img className="avatar__image" src={event.author.avatar.image_url} />
                    </div>
                    <div className="info">
                      <p className="username">{event.author.username}</p>
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