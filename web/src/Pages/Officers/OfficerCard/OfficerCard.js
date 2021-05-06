import React from 'react';
import {
  Card, CardImg, CardText, CardBody,
  CardTitle, CardSubtitle, Button
} from 'reactstrap';
import './OfficerCard.css';
import './Box.css';
import pao from './ACM Officers/pao.jpg';
import lauren from './ACM Officers/lauren.jpg';
import craig from './ACM Officers/craig.jpg';
import kevin from  './ACM Officers/kevin.jpg';
import lindsey from './ACM Officers/lindsey.jpg';

const OfficerCard = (props) => {
  const cardInfo = [
    {image: pao, title: "Kuo-Pao Yang", subtitle: "Faculty Advisor", text: "Kuo-Pao Yang is a professor in the Department of Computer Science at Southeastern Louisiana University. He received his B.S. degree in Computer Science at Tamkang University, Taipei, Taiwan, R.O.C. He earned his M.S. and his Ph.D. degree in Computer Science at the Illinois Institute of Technology. His research interests include Computer Architecture, Programming Languages, and Expert Systems."},
    {image: lauren, title: "Lauren Pace", subtitle: "President", text: "Lauren is a senior Computer Science major. She also interns at the Naval Research Laboratory. After graduating she plans to continue her studies and pursue a Ph.D. "},
    {image: kevin, title: "Kevin Ziebarth", subtitle: "Vice-President", text: "Kevin Ziebarth is a Junior Computer Science Major. He is currently a tutor with the Computer Science Department at Southeastern Louisiana University. Along with being a Peer assistant learning(PAL) who helps teach a software development class.  Outside of school, He enjoys soccer, time with friends and family, working out, cooking, video games, swimming, learning new things, and coding of course."},
    {image: craig, title: "Craig Canepa", subtitle: "Secretary", text: "Craig Canepa is a junior Information Technology major. He is currently a tutor with the Computer Science Department. After receiving his Bachelor's degree, he hopes to get a job. Outside of academics, his interests are playing video games, watching movies, and playing with his pets."},
    {image: lindsey, title: "Lindsey Poche", subtitle: "Treasurer", text: "Lindsey Poche is a Senior Computer Science Major. She is currently a tutor with the Computer Science Department at Southeastern Louisiana University. Along with being the Treasurer for the Association for Computing Machinery, she is also the Treasurer for the American Welding Society at Southeastern Louisiana University. Outside of school, she enjoys swimming and knitting."},
    ];
  
  const renderCard = (card, index) => {
    return (
      <Card key = {index} className= "box">
      <CardImg className="cImage" variant = "top" src={card.image}/>
      <CardBody className="cBody">
        
        <CardTitle className="cTitle" tag="h5">{card.title}</CardTitle>
        <CardSubtitle className="cTitle" tag="h6" className="mb-2 text-muted">{card.subtitle}</CardSubtitle>
        <CardText className="cText">{card.text}</CardText>
        
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