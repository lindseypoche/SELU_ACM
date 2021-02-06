import logo from './logo.svg';
import './App.css';
import NavigationBar from './Components/NavigationBar/NavigationBar';
import { BrowserRouter } from 'react-router-dom';
import { Route, Switch} from 'react-router';
import 'bootstrap/dist/css/bootstrap.min.css';
import Home from './Components/Home/Home.js';
import Login from './Components/LogIn/Login.js';


function App() {
  return (
    <BrowserRouter>
      <NavigationBar/>
      <Switch>
        <Route exact path='/home' component={Home}/>
        <Route exact path='/login' component={Login} />
      </Switch>
    </BrowserRouter>
    
  );
}

export default App;
