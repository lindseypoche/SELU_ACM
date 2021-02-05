import logo from './logo.svg';
import './App.css';
import NavigationBar from './Components/NavigationBar/NavigationBar';
// import SignIn from './Components/SignIn'
import Greet from './Components/Greet'
import Welcome from './Components/Welcome'
import Hello from './Components/Hello'
import Message from './Components/Message'
import Counter from './Components/Counter'
import FunctionClick from './Components/FunctionClick'
import ClassClick from './Components/ClassClick'
import SimpleBreadcrumbs from './Components/SimpleBreadCrumbs'
import EventBind from './Components/EventBind'

function App() {
  return (
    <div className="App">
      <NavigationBar/>
      <Greet name="Tyler" heroName="Superman"> <p> This is children props</p></Greet>
      <Greet name="Diana" heroName="Joker"> 
        <button>Dynamic data that i may not know what it will be, here</button>
      </Greet>
      <Greet name="Bruce" heroName="Joker"/>
      <Welcome name="Tyler" heroName="Superman"> </Welcome>
      <Welcome name="Tyler" heroName="Superman"> </Welcome>
      <Hello/>
      <Message/>
      <Counter> Count </Counter>
      <FunctionClick/>
      <ClassClick/>
      <SimpleBreadcrumbs/>
      <EventBind/>
    </div>
  );
}

export default App;
