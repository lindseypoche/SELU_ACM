import React, { Component } from 'react';
import { Button, Form, FormGroup, Label, Input, FormText } from 'reactstrap';
import './Officers.css';
import OfficerCard from '../OfficerCard/OfficerCard.js';

class Officers extends Component {
    render() {
        return(
            <div className="officersPage">
                <h1>Officers</h1>
                <div className="eventContainer">
                    <OfficerCard/>
                                       
                </div>
            </div>
        )
    }
}
export default Officers;