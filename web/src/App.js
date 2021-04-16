import './App.css';
import NavigationBar from './Components/NavigationBar/NavigationBar';
import { BrowserRouter } from 'react-router-dom';
import { Route, Switch } from 'react-router';
import 'bootstrap/dist/css/bootstrap.min.css';
import Home from './Pages/Home/Home.js';
import Calendar from './Pages/Calendar/CalendarApp.js';
import Officers from './Pages/Officers/Officers.js'
import EventsPage from './Pages/EventsPage/EventsPage.js';
import Resources from './Pages/Resources/Resources.js';
import Footer from './Components/Footer/Footer.js';
import SingleEventPage from './Pages/SingleEventPage/SingleEventPage.js';
import ScrollToTop from './Components/BackTop/BackTop.js';

function App() {

  return (
    <div className="App">
      <BrowserRouter>
        <NavigationBar/>
        <Switch>
          <Route exact path={['/', '/home']} component={Home} />
          <Route path={`/event/:id`} exact strict component={SingleEventPage} />
          <Route exact path='/events' exact strict component={EventsPage} />
          <Route exact path='/calendar' component={Calendar} />
          <Route exact path='/officers' component={Officers} />
          <Route exact path='/resources' component={Resources} />
        </Switch>
        
        {/* Render components in app*/}
      </BrowserRouter>
      <Footer></Footer>
      <ScrollToTop />
    </div>

  );
}

export default App;
