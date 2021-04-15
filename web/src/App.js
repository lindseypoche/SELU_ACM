import './App.css';
import NavigationBar from './Components/NavigationBar/NavigationBar';
import { BrowserRouter } from 'react-router-dom';
import { Route, Switch } from 'react-router';
import 'bootstrap/dist/css/bootstrap.min.css';
import Home from './Components/Home/Home.js';
import Calendar from './Components/Calendar/CalendarApp.js';
import Officers from './Components/Officers/Officers.js'
import EventsPage from './Pages/EventsPage/EventsPage.js';
import Resources from './Components/Resources/Resources';
import Footer from './Components/Footer/Footer.js';
function App() {

  return (
    <div className="App">
      <BrowserRouter>
<<<<<<< HEAD
        <NavigationBar/>
        <Switch>
          <Route exact path='/home' component={Home} />
          <Route exact path='/events' component={EventsPage2} />
=======
        {/* <NavigationBar /> */}
        <Switch>
          <Route exact path='/' component={Home} />
          <Route exact path='/events' component={EventsPage} />
>>>>>>> f7b0f4a (update events page, create event component, etc)
          <Route exact path='/calendar' component={Calendar} />
          <Route exact path='/officers' component={Officers} />
          <Route exact path='/resources' component={Resources} />
          <Route exact path='/singleeventpage' component={SingleEventPage} />
        </Switch>
        
        {/* Render components in app*/}
      </BrowserRouter>
      <Footer></Footer>
      <ScrollToTop />
    </div>

  );
}

export default App;
