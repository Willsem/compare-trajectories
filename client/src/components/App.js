import Map from './Map';
import LoadFile from './LoadFile';
import '../styles/App.css';

function App() {
  return (
    <div class="app">
      <Map class="map" />
      <LoadFile class="files" />
    </div>
  );
}

export default App;
