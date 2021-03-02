import React from 'react';
import {
  Card, CardImg, CardText, CardBody,
  CardTitle, CardSubtitle, Button
} from 'reactstrap';
import './OfficerCard.css';
import Pao from '../image/ACM Officers/pao.resize.jpg'

const OfficerCard = (props) => {
  return (
    <div>
      <Card>
        <CardImg variant = "top" src={Pao} title="Placeholder" />
        <CardBody>
          <CardTitle tag="h5">Officer Name</CardTitle>
          <CardSubtitle tag="h6" className="mb-2 text-muted">Officer Title</CardSubtitle>
          <CardText>Blurb</CardText>
          <Button>Read More</Button>
        </CardBody>
      </Card>
      
    </div>
  );
};

export default OfficerCard;