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

/*
const Officers = () => {

    const [officers, setOfficers] = useState([]);
    const [officersIsLoaded, setOfficersIsLoaded] = useState(false)
    const [officersError, setOfficersError] = useState(false)

    useEffect(() => {
        getEvent();
    }, [])

    const getEvent = () => {
       axios.get(`http://localhost:8081/api/members/officers`)
           .then(((response) => {  
               setOfficers(response.data);
               setOfficersIsLoaded(true);
          }))
          .catch(error => {
            console.log(error);
            setOfficersError(true);
          })
      }

      if (!officersIsLoaded) {
        return <div>Loading...</div>;
      }

      if (officersError) {
        return <div>Error fetching officer data</div>;
      }

      return (
        <div className="officersPage">
            <div className="officersPara">
                <h1>Officers</h1>
                <div className="officersContainer">
                    {
                      officersError ? (
                        'Error fetching officer data'
                      ) : (
                        <OfficerCard officers={officers} />
                      )
                    }
                </div>
            </div>
            <div>
                
            </div>
        </div>
    )
}
*/