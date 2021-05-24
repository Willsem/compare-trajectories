import { MapContainer, TileLayer, Polyline, LayersControl } from 'react-leaflet';
import { filtering } from '../api/filtering'
import 'leaflet/dist/leaflet.css';
import '../styles/Map.css';

function Map() {
  let trajectory = {};
  let loaded = false;

  let position = [0, 0];
  let zoom = 13;

  let polyline = [];

  if (loaded) {
    trajectory = filtering(trajectory);

    for (let i = 0; i < trajectory.long.length; ++i) {
      polyline[i] = [trajectory.long[i], trajectory.lat[i]];

      position[0] += polyline[i][0];
      position[1] += polyline[i][1];
    }

    position[0] /= polyline.length;
    position[1] /= polyline.length;
  } else {
    zoom = 2;
  }

  const optionsRed = {color: "red"};
  return (
    <MapContainer className='map-container' center={position} zoom={zoom} scrollWheelZoom={true}>

      <LayersControl position="topright">
        <LayersControl.Overlay checked name="Good Trajectory">
          <Polyline pathOptions={optionsRed} positions={polyline} />
        </LayersControl.Overlay>
        <LayersControl.Overlay checked name="Compared Trajectory">
          <Polyline pathOptions={optionsRed} positions={polyline} />
        </LayersControl.Overlay>
      </LayersControl>

      <TileLayer
        attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />
    </MapContainer>
  );
}

export default Map;
