// import { FaTimes } from 'react-icons/fa'
import './Event.css'


// getTime(unixTime) {
//     var t = new Date();
//     t.setSeconds( unixTime );
//     unixTime = t.format("dd.mm.yyyy hh:MM:ss");
// }


const Event = ({ event }) => {
    return (
        <div className="card">

        { event.attachments != null ? (
            <>
                <img src={event.attachments.url} />
                <div className="container">
                  <h4><b>{event.author.username}</b></h4>
                  <p>{event.content.substring(0,15)}</p>
                </div>
            </>
            ) : (
            <>
                <img src='https://www.startpage.com/av/proxy-image?piurl=https%3A%2F%2Fi.insider.com%2F5b9fdb115c5e5236008b6334%3Fwidth%3D750%26format%3Djpeg%26auto%3Dwebp&sp=1615235049Tfe91e30835fc44f1702dc7cb4638dce12341708b38942fad76ba5df68a14e3fe' />
                <div className="container">
                  <h4><b>{event.author.username}</b></h4>
                  <p>{event.content.substring(0,15)}</p>
                </div>
            </>
            )
        }
        </div>
    )
}

export default Event
