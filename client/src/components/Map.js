import { MapContainer, TileLayer, Polyline, LayersControl } from 'react-leaflet';
import { filtering } from '../api/filtering'
import ChangeMapView from './ChangeMapView';
import 'leaflet/dist/leaflet.css';
import '../styles/Map.css';

function convertToPolyline(trajectory) {
  let polyline = [];

  if (trajectory && !trajectory.error) {
    for (let i = 0; i < trajectory.long.length; ++i) {
      polyline[i] = [trajectory.long[i], trajectory.lat[i]];
    }
  }

  return polyline;
}

const optionsRed = {color: "red"};
const optionsGreen = {color: "green"};

function Map({ perfectTrajectory, comparedTrajectory, position, zoom }) {
  let filteredPerfectTrajectory = filtering(perfectTrajectory.gps);
  let filteredComparedTrajectory = comparedTrajectory.gps // filtering(comparedTrajectory.gps);

  let perfectPolyline = convertToPolyline(filteredPerfectTrajectory);
  let comparedPolyline = convertToPolyline(filteredComparedTrajectory);

  return (
    <MapContainer className='map-container' center={position} zoom={zoom} scrollWheelZoom={true}>
      <ChangeMapView position={position} zoom={zoom} />

      <LayersControl position="topright">
        <LayersControl.Overlay checked name="Good Trajectory">
          <Polyline pathOptions={optionsGreen} positions={perfectPolyline} />
        </LayersControl.Overlay>
        <LayersControl.Overlay checked name="Compared Trajectory">
          <Polyline pathOptions={optionsRed} positions={comparedPolyline} />
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
