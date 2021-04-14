import './App.css';
import NavigationBar from './Components/NavigationBar/NavigationBar';
import { BrowserRouter } from 'react-router-dom';
import { Route, Switch } from 'react-router';
import 'bootstrap/dist/css/bootstrap.min.css';
import Home from './Pages/Home/Home.js';
import Calendar from './Pages/Calendar/CalendarApp.js';
import Officers from './Pages/Officers/Officers.js'
import EventsPage2 from './Pages/EventsPage/EventsPage2.js';
import Resources from './Pages/Resources/Resources.js';
// import Footer from './Unused Components/Footer/Footer.js';
import SingleEventPage from './Pages/SingleEventPage/SingleEventPage.js';
import ScrollToTop from './Components/BackTop/BackTop.js';

function App() {

  return (
    <div className="App">
      <BrowserRouter>
        <NavigationBar/>
        <Switch>
          <Route exact path='/home' component={Home} />
          <Route exact path='/events' component={EventsPage2} />
          <Route exact path='/calendar' component={Calendar} />
          <Route exact path='/officers' component={Officers} />
          <Route exact path='/resources' component={Resources} />
          <Route exact path='/singleeventpage' component={SingleEventPage} />
        </Switch>
        <ScrollToTop />
        {/* Render components in app*/}
      </BrowserRouter>

    </div>

  );
}

export default App;
