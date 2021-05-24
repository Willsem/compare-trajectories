import DropZoneFile from './DropZoneFile';
import '../styles/LoadFile.css'

function LoadFiles() {
  return (
    <div class="load-container">
      <section class="container">
        <h4>Perfect trajectory</h4>
        <DropZoneFile />
      </section>

      <section class="container">
        <h4>Compared trajectory</h4>
        <DropZoneFile />
      </section>
    </div>
  );
}

export default LoadFiles;
