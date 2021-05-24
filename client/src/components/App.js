import { Component } from 'react';
import Map from './Map';
import LoadFile from './LoadFile';
import '../styles/App.css';

class App extends Component {
  render() {
    return (
      <div class="app">
        <Map class="map" />
        <LoadFile class="files" />
      </div>
    );
  }
}

export default App;
