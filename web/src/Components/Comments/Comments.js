import axios from 'axios';
import React,{useState, useEffect} from 'react';


const Comments = (id) =>{

    const [comments, setComments] = useState([]);
    const [commentsLoaded, setCommentsLoaded] = useState(false);

    const GetComments = () => {
        axios.get(`http://localhost/8081/events/${id}/comments`)
        .then((res)=>{
            setComments(res.data);
            setCommentsLoaded(true);
        }).catch((err)=>{
            console.log(err);
        });

        if(!commentsLoaded){
            setCommentsLoaded(false);
           return( <div>...Loading</div>);
        }
    }

        useEffect(()=>{
            GetComments();
        }, []);

        return(
            <div className ="outerWrapper">
                <div className="Avatar">
                    <div className = "Name">
                        <div className = "CommentBody">
                            <h2>HI</h2>
                        </div>
                    </div>
                </div>  
            </div>
        );

    }

export default Comments;