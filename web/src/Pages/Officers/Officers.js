import React, { Component } from 'react';
import './Officers.css';
import OfficerCard from './OfficerCard/OfficerCard.js';

class Officers extends Component {
    render() {
        return (
            <div className="officersPage">
                <div className="officersPara">
                    <h1>Officers</h1>
                    <div className="officersContainer">
                        <OfficerCard />
                    </div>
                </div>
                <div>
                    
                </div>
            </div>
        )
    }
}
export default Officers;