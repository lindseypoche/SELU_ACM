import './App.css';
import NavigationBar from './Components/NavigationBar/NavigationBar';
import { BrowserRouter } from 'react-router-dom';
import { Route, Switch } from 'react-router';
import 'bootstrap/dist/css/bootstrap.min.css';
import Home from './Components/Home/Home.js';
import Calendar from './Components/Calendar/Calendar.js';
import Officers from './Components/Officers/Officers.js'
import EventsPage2 from './Pages/EventsPage/EventsPage2.js';
import Resources from './Components/Resources/Resources';
import Footer from './Components/Footer/Footer.js';


function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <NavigationBar />
        <Switch>
          <Route exact path='/' component={Home} />
          <Route exact path='/events' component={EventsPage2} />
          <Route exact path='/calendar' component={Calendar} />
          <Route exact path='/officers' component={Officers} />
          <Route exact path='/resources' component={Resources} />
        </Switch>
        
        {/* Render components in app*/}
      </BrowserRouter>

    </div>
    
  );
}

export default App;
