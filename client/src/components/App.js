import { Component } from 'react';
import Map from './Map';
import '../styles/App.css';

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
