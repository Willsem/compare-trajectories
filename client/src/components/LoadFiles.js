import DropZoneFile from './DropZoneFile';
import '../styles/LoadFile.css'

function LoadFiles() {
  return (
    <div className="load-container">
      <section className="container">
        <h4>Perfect trajectory</h4>
        <DropZoneFile />
      </section>

      <section className="container">
        <h4>Compared trajectory</h4>
        <DropZoneFile />
      </section>
    </div>
  );
}

export default LoadFiles;
