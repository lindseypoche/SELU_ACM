import axios from 'axios';
import {GET_EVENTS} from '../types';


const initialState = {
    events:[],
    loading:true
}

export default function(state = initialState, action){
    // if(typeof state === 'undefined'){
    //     return state;
    // }

    if(action.type === 'GET_EVENTS'){
        return {
            ...state,
            events: action.payload,
            loading : false
        }
    }else{
        return state;
    }
}