import logo from './logo.svg';
import './App.css';
import NavigationBar from './Components/NavigationBar/NavigationBar';
import { BrowserRouter } from 'react-router-dom';
import { Route, Switch } from 'react-router';
import 'bootstrap/dist/css/bootstrap.min.css';
import Home from './Components/Home/Home.js';
import Login from './Components/LogIn/Login.js';
import VideoPlayback from './Components/video/VideoPlayback.js'

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <NavigationBar />
        <VideoPlayback />
        <Switch>
          <Route exact path='/' component={Home} />
          <Route exact path='/login' component={Login} />
        </Switch>
      </BrowserRouter>
    </div>
  );
}

export default App;
