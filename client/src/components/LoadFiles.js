import DropZoneFile from './DropZoneFile';
import '../styles/LoadFile.css'

function LoadFiles() {
  const loadPerfectFile = (file) => {
    console.log('perfect');
    console.log(file);
  };

  const loadComparedFile = (file) => {
    console.log('compared');
    console.log(file);
  };

  return (
    <div className="load-container">
      <section className="container">
        <h4>Perfect trajectory</h4>
        <DropZoneFile loadFileCallback={loadPerfectFile} />
      </section>

      <section className="container">
        <h4>Compared trajectory</h4>
        <DropZoneFile loadFileCallback={loadComparedFile} />
      </section>
    </div>
  );
}

export default LoadFiles;
