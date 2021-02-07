import logo from './logo.svg';
import './App.css';
import NavigationBar from './Components/NavigationBar/NavigationBar';
import { BrowserRouter } from 'react-router-dom';
import { Route, Switch} from 'react-router';
import 'bootstrap/dist/css/bootstrap.min.css';
import Home from './Components/Home/Home.js';
import Login from './Components/LogIn/Login.js';
import Members from './Components/Members/Members.js';
import Events from './Components/Events/Events.js';
import About from './Components/About/About.js';
import Join from './Components/Join/Join.js';


function App() {
  return (
    <BrowserRouter>
      <NavigationBar/>
      <Switch>
        <Route exact path='/home' component={Home}/>
        <Route exact path='/events' component={Events}/>
        <Route exact path='/about' component={About}/>
        <Route exact path='/join' component={Join}/>
        <Route exact path='/login' component={Login}/>
        <Route exact path='/members' component={Members}/>
        
      </Switch>
    </BrowserRouter>
    
  );
}

export default App;
