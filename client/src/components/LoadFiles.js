import DropZoneFile from './DropZoneFile';
import '../styles/LoadFile.css'

function LoadFiles({ setPerfectTrajectory, setComparedTrajectory }) {
  const loadPerfectFile = (file) => {
    setPerfectTrajectory(file);
  };

  const loadComparedFile = (file) => {
    setComparedTrajectory(file);
  };

  return (
    <div className="load-container">
      <section className="container">
        <h4>Perfect trajectory</h4>
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
