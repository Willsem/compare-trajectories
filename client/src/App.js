import { Component } from 'react';
import Map from './components/Map/Map';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props)
  }

  render() {
    return (
      <div>
        <Map />
      </div>
    );
  }
}

export default App;
