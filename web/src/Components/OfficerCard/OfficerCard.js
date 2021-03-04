import React from 'react';
import {
  Card, CardImg, CardText, CardBody,
  CardTitle, CardSubtitle, Button
} from 'reactstrap';
import './OfficerCard.css';
import './Box.css'



const OfficerCard = (props) => {
  const cardInfo = [
    {image: "", title: "Kuo-Pao Yang", subtitle: "Faculty Advisor", text: "Kuo-pao Yang is a professor in the Department of Computer Science at Southeastern Louisiana University.He received his B.S. degree in Computer Science at Tamkang University, Taipei, Taiwan, R.O.C. He earned his M.S. and his Ph.D. degree in Computer Science at Illinois Institute of Technology. His research interests include Computer Architecture, Programming Languages, and Expert Systems."},
    {image: "", title: "Lauren Pace", subtitle: "President", text: ""},
    {image: "",title: "Kevin Ziebarth", subtitle: "Vice-President", text: ""},
    {image: "",title: "Craig Canepa", subtitle: "Secretary", text: "Craig Canepa is a junior Information Technology major. He is currently a tutor with the Computer Science Department. After receiving his Bachelors degree, he hopes to get a job. Outside of academics, his interests are playing video games, watching movies, and playing with his pets."},
    {image: "",title: "Lindsey Poche", subtitle: "Treasurer", text: "Lindsey Poche is a Senior Computer Science Major.She is currently a tutor with the Computer Science Department as Southeastern Lousisiana University. Along with being the Treasurer for the Association for Computing Machinery, she is also the Treasurer for American Welding Society at Southeastern Louisiana University. Outside of school she enjoys swimming and knitting."},
    ];
  
  const renderCard = (card, index) => {
    return (
      <Card style = {{width: "18rem "}} key = {index} className= "box">
      <CardImg variant = "top" src={card.image}/>
      <CardBody>
        
        <CardTitle tag="h5">{card.title}</CardTitle>
        <CardSubtitle tag="h6" className="mb-2 text-muted">{card.subtitle}</CardSubtitle>
        <CardText>{card.text}</CardText>
        <Button>Read More</Button>
      </CardBody>
    </Card>
    )
  };
  return (
    <div className = "grid">
        {cardInfo.map(renderCard)}    
    </div>
  );
};

export default OfficerCard;