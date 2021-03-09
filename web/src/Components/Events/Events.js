import React, { useState, useEffect } from "react";
import { Button, Form, FormGroup, Label, Input, FormText } from "reactstrap";
import "./Events.css";
import BlogCard from "../Card2/BlogCard.js";
import axios from "axios";

function Events(props) {
    // var Urlid = props.match.params.id;
    const [someHook, setSomeHook] = useState([]);
    useEffect(() => {

        axios
        // .get(`http://localhost:8080/blogs/${props.match.params.id}`)
            .get(`http://localhost:8080/blogs`)
            .then((result) => {
                setSomeHook(result.data);
            })
            .catch((err) => {
                console.log(err);
            });
    }, [])

    var cards = [];

    for (var i in someHook){
        console.log(someHook)
        cards.push(<BlogCard 
            dataobj={someHook[i]}
        />)
    }

    return (
    <div className='eventContainer'>
         { cards } 
    </div>
    );
}

export default Events;
