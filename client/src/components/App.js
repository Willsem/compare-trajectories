import { useState } from 'react';
import Map from './Map';
import LoadFiles from './LoadFiles';
import '../styles/App.css';

function App() {
  const [perfectTrajectory, setPerfectTrajectory] = useState(0);
  const [comparedTrajectory, setComparedTrajectory] = useState(0);

  return (
    <div className="app">
      <Map className="map"
        perfectTrajectory={perfectTrajectory}
        comparedTrajectory={comparedTrajectory} />

      <LoadFiles className="files"
        setPerfectTrajectory={setPerfectTrajectory}
        setComparedTrajectory={setComparedTrajectory} />
    </div>
  );
}

export default App;
