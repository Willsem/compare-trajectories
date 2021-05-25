import { MapContainer, TileLayer, Polyline, LayersControl } from 'react-leaflet';
import { filtering } from '../api/filtering'
import 'leaflet/dist/leaflet.css';
import '../styles/Map.css';

function convertToPolyline(trajectory) {
  let polyline = [];

  if (!trajectory.error) {
    for (let i = 0; i < trajectory.long.length; ++i) {
      polyline[i] = [trajectory.long[i], trajectory.lat[i]];
    }
  }

  return polyline;
}

function getPosition(polyline1, polyline2) {
  if (polyline1.length + polyline2.length === 0) {
    return [[0, 0], 2];
  }

  return [[
    ((polyline1.length ? polyline1.reduce((s, c) => s + c[0]) : 0) +
    (polyline2.lenght ? polyline2.reduce((s, c) => s + c[0]) : 0)) /
    (polyline1.length + polyline2.length),

    ((polyline1.lenght ? polyline1.reduce((s, c) => s + c[1]) : 0) +
    (polyline2.lenght ? polyline2.reduce((s, c) => s + c[1]) : 0)) /
    (polyline1.length + polyline2.length),
  ], 13];
}

const optionsRed = {color: "red"};
const optionsGreen = {color: "green"};

function Map({ perfectTrajectory, comparedTrajectory }) {
  let filteredPerfectTrajectory = filtering(perfectTrajectory.gps);
  let filteredComparedTrajectory = filtering(comparedTrajectory.gps);

  let perfectPolyline = convertToPolyline(filteredPerfectTrajectory);
  let comparedPolyline = convertToPolyline(filteredComparedTrajectory);

  let [position, zoom] = getPosition(perfectPolyline, comparedPolyline);

  return (
    <MapContainer className='map-container' center={position} zoom={zoom} scrollWheelZoom={true}>

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
