import { useMap } from 'react-leaflet';

function ChangeMapView({ position, zoom }) {
  const map = useMap();
  map.flyTo(position, zoom, {
    duration: 2
  });

  return null;
}

export default ChangeMapView;
