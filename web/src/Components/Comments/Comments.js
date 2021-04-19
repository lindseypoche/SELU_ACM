import axios from 'axios';
import React,{useState, useEffect} from 'react';
import "./Comments.css";


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
            <div className ="Comments_Wrapper">
                <span>Comments</span>
                <div className="Avatar">
                    <div className = "Name">
                        <div className = "CommentBody">
                        </div>
                    </div>
                </div>  
            </div>
        );

    }

export default Comments;