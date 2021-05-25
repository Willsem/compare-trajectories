import { useState } from 'react';
import Map from './Map';
import LoadFiles from './LoadFiles';
import '../styles/App.css';

function App() {
  const [perfectTrajectory, setPerfectTrajectory] = useState(0);
  const [comparedTrajectory, setComparedTrajectory] = useState(0);

  const [position, setPosition] = useState([0, 0]);
  const [zoom, setZoom] = useState(2);

  return (
    <div className="app">
      <Map className="map"
        perfectTrajectory={perfectTrajectory}
        comparedTrajectory={comparedTrajectory}
        position={position}
        zoom={zoom} />

      <LoadFiles className="files"
        setPerfectTrajectory={setPerfectTrajectory}
        setComparedTrajectory={setComparedTrajectory}
        setPosition={setPosition}
        setZoom={setZoom} />
    </div>
  );
}

export default App;
