import DropZoneFile from './DropZoneFile';
import '../styles/LoadFile.css'

function getPosition(trajectory) {
  if (!trajectory || !trajectory.long || !trajectory.lat) {
    return [[0, 0], 2];
  }

  return [[
    trajectory.long.reduce((s, c) => s + c) / trajectory.long.length,
    trajectory.lat.reduce((s, c) => s + c) / trajectory.lat.length,
  ], 16];
}

function LoadFiles({ setPerfectTrajectory, setComparedTrajectory, setPosition, setZoom }) {
  const updatePos = (trajectory) => {
    let [pos, z] = getPosition(trajectory.gps);
    setPosition(pos);
    setZoom(z);
  };

  const loadPerfectFile = (file) => {
    updatePos(file);
    setPerfectTrajectory(file);
  };

  const loadComparedFile = (file) => {
    updatePos(file);
    setComparedTrajectory(file);
  };

  return (
    <div className="load-container">
      <section className="container">
        <h4>Reference trajectory</h4>
        <DropZoneFile loadFileCallback={loadPerfectFile} fieldName="perfect" />
      </section>

      <section className="container">
        <h4>Compared trajectory</h4>
        <DropZoneFile loadFileCallback={loadComparedFile} fieldName="compared" />
      </section>
    </div>
  );
}

export default LoadFiles;
