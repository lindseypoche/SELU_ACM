import axios from 'axios';
import React,{useState, useEffect} from 'react';
import "./Comments.css";
import { toDateFormat, isExpiring, getRemainingTime } from "../../Utils/timing.js"

const Comments = (match) => {

    const [comments, setComments] = useState([]);
    const [commentsLoaded, setCommentsLoaded] = useState(false);

    useEffect(() => {
        GetComments();
    }, []);

    const GetComments = () => {
        axios.get(`http://localhost:8081/api/events/${match.eventID}/comments`)
            .then((res) => {
                setComments(res.data);
                setCommentsLoaded(true);
            })
            .catch(error => console.log(error))
    }

    if(!commentsLoaded){
       return ( <div>Loading...</div>);
    }

    return(
        <div className="main">
            <div className="comments-header">Comments</div>
            <div className ="comments-container">
            {
                comments.map((comment) => (
                    <div className="comment-wrapper" key={comment.id}>
                        <article className="avatar">
                            <img src={comment.author.avatar.image_url} />
                        </article>
                        <article className="text-wrapper">
                            <div className="username-date-wrapper">
                                <span className="username">{comment.author.username}</span>
                                <span className="datetime">{toDateFormat(comment.timestamp)}</span>
                            </div>
                            <p className="content-body">{comment.content}</p>
                        </article>
                    </div>
                ))
            }
            </div>
        </div>
    );

    }

export default Comments;