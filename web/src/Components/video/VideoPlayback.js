import React from 'react'
import finalcut from '../video/finalcut0000-0902.mp4'
import "VideoP";

const VideoPlayback = () => {
    return (
        <div className  style={{height: '0px', width:'0px', overflow: 'visible', float:'left'}}>
            <video
                autoPlay
                loop
                style= {{
                    
                    position: "relative",
                    //width: '100vw',
                    // left: '50%',
                    zIndex: '98',
                    top: '200px',
                    height: '50em',
                    float: 'left',
                    objectFit: 'cover',
                    // transform: 'translate(-50%, -50%)',
                    zIndex: '-1',
                    filter:"grayscale(30%)",
                }}
                >
                    <source src= {finalcut} type='video/mp4'/>
                </video>
        </div>
    )
}

export default VideoPlayback
