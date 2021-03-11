import slugo from "./slugo.png"

import CardColumns from 'react-bootstrap/CardColumns';
import Card from 'react-bootstrap/Card';

const Events = ({ events }) => {

    return (
        <CardColumns>
        {
          (events.length > 0 ? (
          events.map((event) => (
          <Card key={event.id} className="p-3">
              { event.attachments != null ? 
                    (
                        <Card.Img variant="top" src={event.attachments.url} />
                    ) : (
                        // <Card.Img variant="top" src={slugo} />
                        <div>No image available</div>
                    )
              }
            <Card.Body>
                <Card.Title>{event.content.substring(0, 15)}</Card.Title>
                <Card.Text style={{textAlign: "left"}}>
                    {event.content.substring(0, 200)}
                </Card.Text>
              </Card.Body>
              <Card.Footer>
                <img style={{marginRight: "10px", height: "50px", width: "50px", borderRadius: "50%", textAlign: "left"}} className='discord-avatar' src={event.author.avatar.image_url} />
                <small className="text-muted">edited at {event.timestamp}</small>
              </Card.Footer>
            </Card>
          ))
            ) : (
            'No events to show'
            )
          )
        }
        </CardColumns>
    )
}

export default Events;
