import React from 'react'
import finalcut from '../video/finalcut0000-0902.mp4'

const VideoPlayback = () => {
    return (
        <div style={{height: '0px', width:'0px', overflow: 'visible', float:'left'}}>
            <video
                autoPlay
                loop
                style= {{
                    
                    position: "fixed",
                    width: '100vw',
                    
                    // left: '50%',
                    // top: '50%',
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


// var NewComponent = React.createClass({
//     render: function() {
//       return (
//         <div>
//           <header className="video-header">
//             <video src="https://css-tricks-post-videos.s3.us-east-1.amazonaws.com/Island%20-%204141.mp4" autoPlay loop playsInline muted />
//             <div className="viewport-header">
//               <h1>
//                 Explore
//                 <span>Montana</span>
//               </h1>
//             </div>
//           </header>
//           <main>
//             [[[https://codepen.io/chriscoyier/pen/VbqzER]]]
//           </main>
//           <div id="circle" />
//         </div>
//       );
//     }
//   });



// var viewportHeader = document.querySelector(".viewport-header");

// document.body.addEventListener("scroll", function(event) {
//   var opacity = (document.body.offsetHeight - document.body.scrollTop) / document.body.offsetHeight;
//   var scale = (document.body.offsetHeight - document.body.scrollTop) / document.body.offsetHeight;
//   document.documentElement.style.setProperty('--headerOpacity', opacity);
//   document.documentElement.style.setProperty('--headerScale', scale);
// });

