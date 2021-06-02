import {
  MapContainer,
  TileLayer,
  Polyline,
  LayersControl,
  LayerGroup,
  Tooltip,
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
  if (backlog > 0.1) {
    return 'green';
  }

  if (backlog < -0.1) {
    return 'red';
  }

  return 'yellow';
}

function difference(array) {
  return array[0] - array[array.length - 1];
}

function pointDifference(array) {
  return {
    'x': array[0].x - array[array.length - 1].x,
    'y': array[0].y - array[array.length - 1].y,
    'z': array[0].z - array[array.length - 1].z,
  };
}

const optionsPerfect = {color: 'black'};
const optionsCompared = {color: 'grey'};

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
      const options = {color: convertBacklogToColor(difference(backlog))};
      const polyline = convertToPolyline(compareResult[i])
      comparedTrajectoryElement.push({'option': options, 'positions': polyline, 'diff': {
        'backlog': difference(backlog),
        'dlong': difference(compareResult[i].dlong),
        'dlat': difference(compareResult[i].dlat),
        'dacc': pointDifference(compareResult[i].dacc),
        'dgyro': pointDifference(compareResult[i].dgyro),
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
        <LayersControl.Overlay checked name="Эталонная траектория">
          <Polyline pathOptions={optionsPerfect} positions={perfectPolyline} />
        </LayersControl.Overlay>
        <LayersControl.Overlay checked name="Сравниваемая траектория">
          <LayerGroup>
            {comparedTrajectoryElement.map(item =>
              <div>
                <Polyline pathOptions={item.option} positions={item.positions}>
                  <Tooltip sticky>
                    <pre>
                      {'Отставание: ' + item.diff.backlog + '\n\n' +
                      'Разница широты: ' + item.diff.dlong + '\n' +
                      'Разница долготы: ' + item.diff.dlat + '\n\n' +
                      'Разница акселерометра:\n' +
                      'x: ' + item.diff.dacc.x + '\n' +
                      'y: ' + item.diff.dacc.y + '\n' +
                      'z: ' + item.diff.dacc.z + '\n\n' +
                      'Разница гироскопа:\n' +
                      'x: ' + item.diff.dgyro.x + '\n' +
                      'y: ' + item.diff.dgyro.y + '\n' +
                      'z: ' + item.diff.dgyro.z}
                    </pre>
                  </Tooltip>
                </Polyline>
              </div>
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
