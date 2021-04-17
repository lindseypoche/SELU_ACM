import {GET_EVENTS, EVENTS_ERROR} from '../types';
import axios from 'axios';

export default getEvents = ()  => async dispatch  => {

    try{
        const res = await axios.get(`http://jsonplaceholder.typicode.com/users`)
        dispatch( {
            type: GET_EVENTS,
            payload: res.data
        })
    }
    catch(e){
        dispatch( {
            type: EVENTS_ERROR,
            payload: console.log(e),
        })
    }

}