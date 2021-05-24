import { Component } from 'react';
import { MapContainer, TileLayer, Polyline } from 'react-leaflet';
import { filtering } from '../api/filtering'
import 'leaflet/dist/leaflet.css';
import '../styles/Map.css';
import data from './gps13.json';

class Map extends Component {
  render() {
    const optionsRed = {color: "red"};
    let trajectory = {
      'date': [],
      'long': [],
      'lat': [],
    };
    data.forEach((element) => {
      trajectory.long[trajectory.long.length] = element.long;
      trajectory.lat[trajectory.lat.length] = element.lat;
    });

    trajectory = filtering(trajectory);

    let position = [0, 0];
    let polyline = [];

    for (let i = 0; i < trajectory.long.length; ++i) {
      polyline[i] = [trajectory.long[i], trajectory.lat[i]];

      position[0] += polyline[i][0];
      position[1] += polyline[i][1];
    }

    position[0] /= polyline.length;
    position[1] /= polyline.length;

    return (
      <MapContainer className='map-container' center={position} zoom={13} scrollWheelZoom={true}>
        <TileLayer
          attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />
        <Polyline pathOptions={optionsRed} positions={polyline} />
      </MapContainer>
    );
  }
}

export default Map;
