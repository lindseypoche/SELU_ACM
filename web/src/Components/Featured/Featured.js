// import React, { useState, useEffect } from "react";
// import axios from "axios";
import './Featured.css'

const Featured = ({ featured }) => {
 
    return (
        <>
        <div className='featured'>

            <p>{featured.channel_id}</p>
            <img src={featured.message.attachments.url} />
            {/* <div className='content'>
                <p>{featured.message.content}</p>
            </div> */}
        </div>
        < br/><br/><br/>
        </>
    )
}

export default Featured
