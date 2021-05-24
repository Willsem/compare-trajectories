import Map from './Map';
import LoadFiles from './LoadFiles';
import '../styles/App.css';

function App() {
  return (
    <div class="app">
      <Map class="map" />
      <LoadFiles class="files" />
    </div>
  );
}

export default App;
