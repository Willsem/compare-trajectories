import {
  MapContainer,
  TileLayer,
  Polyline,
  LayersControl,
  LayerGroup,
  Tooltip
} from 'react-leaflet';

import { filtering } from '../api/filtering'
import { compare } from '../api/compare'
import ChangeMapView from './ChangeMapView';
import 'leaflet/dist/leaflet.css';
import '../styles/Map.css';

function convertToPolyline(trajectory) {
  let polyline = [];

  if (trajectory && !trajectory.error) {
    for (let i = 0; i < trajectory.long.length; ++i) { polyline[i] = [trajectory.long[i], trajectory.lat[i]];
    }
  }

  return polyline;
}

function convertBacklogToColor(backlog) {
  if (backlog > 0.3) {
    return 'green';
  }

  if (backlog < -0.3) {
    return 'red';
  }

  return 'yellow';
}

const optionsPerfect = {color: 'black'};
const optionsCompared = {color: 'grey'}

function Map({ perfectTrajectory, comparedTrajectory, position, zoom }) {
  if (!perfectTrajectory.filtered) {
    perfectTrajectory.gps = filtering(perfectTrajectory.gps);
    perfectTrajectory.filtered = true;
  }

  if (!comparedTrajectory.filtered) {
    comparedTrajectory.gps = filtering(comparedTrajectory.gps);
    comparedTrajectory.filtered = true;
  }

  const perfectPolyline = convertToPolyline(perfectTrajectory.gps);
  const comparedPolyline = convertToPolyline(comparedTrajectory.gps);
  const compareResult = compare(perfectTrajectory, comparedTrajectory);

  let comparedTrajectoryElement = [];
  if (compareResult && !compareResult.error) {
    for (let i = 0; i < compareResult.length; ++i) {
      const backlog = compareResult[i].backlog;
      const options = {color: convertBacklogToColor(backlog[0] - backlog[backlog.length - 1])};
      const polyline = convertToPolyline(compareResult[i])
      console.log(compareResult[i]);
      comparedTrajectoryElement.push({'option': options, 'positions': polyline, 'diff': {
        'backlog': backlog,
        'dlong': compareResult[i].dlong[0],
        'dlat': compareResult[i].dlat[0],
        'dacc': compareResult[i].dacc[0],
        'dgyro': compareResult[i].dgyro[0],
      }});
    }
  } else {
    comparedTrajectoryElement.push({'option': optionsCompared, 'positions': comparedPolyline, 'diff': {
        'backlog': '???',
        'dlong': '???',
        'dlat': '???',
        'dacc': '???',
        'dgyro': '???',
    }});
  }

  return (
    <MapContainer className='map-container' center={position} zoom={zoom} scrollWheelZoom={true}>
      <ChangeMapView position={position} zoom={zoom} />

      <LayersControl position="topright">
        <LayersControl.Overlay checked name="Reference Trajectory">
          <Polyline pathOptions={optionsPerfect} positions={perfectPolyline} />
        </LayersControl.Overlay>
        <LayersControl.Overlay checked name="Compared Trajectory">
          <LayerGroup>
            {comparedTrajectoryElement.map(item =>
              <Polyline pathOptions={item.option} positions={item.positions}>
              // TODO: fix
                <Tooltip sticky>
                  <pre>
                    {'Backlog: ' + item.diff.backlog + '\n' +
                    'Delta longitude: ' + item.diff.dlong + '\n' +
                    'Delta latitude: ' + item.diff.dlat + '\n' +
                    'Delta accelerometer:\n' +
                    'x: ' + item.diff.dacc.x + '\n' +
                    'y: ' + item.diff.dacc.y + '\n' +
                    'z: ' + item.diff.dacc.z + '\n' +
                    'Delta gyroscope:\n' +
                    'x: ' + item.diff.dgyro.x + '\n' +
                    'y: ' + item.diff.dgyro.y + '\n' +
                    'z: ' + item.diff.dgyro.z}
                  </pre>
                </Tooltip>
              </Polyline>
            )}
          </LayerGroup>
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
