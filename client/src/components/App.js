import { Component } from 'react';
import Map from './Map';
import LoadFile from './LoadFile';
import '../styles/App.css';

class App extends Component {
  render() {
    return (
      <div>
        <Map />
        <LoadFile />
      </div>
    );
  }
}

export default App;
