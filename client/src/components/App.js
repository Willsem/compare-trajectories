import Map from './Map';
import LoadFiles from './LoadFiles';
import '../styles/App.css';

function App() {
  return (
    <div className="app">
      <Map className="map" />
      <LoadFiles className="files" />
    </div>
  );
}

export default App;
