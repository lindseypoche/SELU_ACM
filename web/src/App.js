import './App.css';
import NavigationBar from './Components/NavigationBar/NavigationBar';
import { BrowserRouter } from 'react-router-dom';
import { Route, Switch } from 'react-router';
import 'bootstrap/dist/css/bootstrap.min.css';
import Home from './Components/Home/Home.js';
import Login from './Components/LogIn/Login.js';
import Calendar from './Components/Calendar/Calendar.js';
import Officers from './Components/Officers/Officers.js'
import Events from './Components/Events/Events';




function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <NavigationBar />
        <Switch>
          <Route exact path='/' component={Home} />
          <Route exact path='/login' component={Login} />
          <Route exact path='/events' component={Events} />
          <Route exact path='/calendar' component={Calendar} />
          <Route exact path='/officers' component={Officers} />

        </Switch>
      </BrowserRouter>
    </div>
    
  );
}

export default App;
