"use client";
import { useLoadScript, GoogleMap, Marker, Polyline } from "@react-google-maps/api";
import { useRef, useCallback } from "react";
import { Ride } from "@/types/Ride";

const mapContainerStyle = { width: "100%", height: "100%" };

interface RideMapProps {
  rides: Ride[];
  selectedRideId: string | null;
  onPolylineClick: () => void;
}

export default function RideMap({
  rides,
  selectedRideId,
  onPolylineClick,
}: RideMapProps) {
  const { isLoaded, loadError } = useLoadScript({
    googleMapsApiKey: process.env.NEXT_PUBLIC_GMAPS_KEY!,
  });

  const mapRef = useRef<google.maps.Map | null>(null);
  const onMapLoad = useCallback((map: google.maps.Map) => {
    mapRef.current = map;
  }, []);

  if (loadError) return <div>Error loading map</div>;
  if (!isLoaded) return <div>Loading map…</div>;

  const selectedRide = rides.find(r => r.id === selectedRideId);

  const defaultCenter = rides[0]?.pickup ?? { lat: 0, lng: 0 };
  const center = selectedRide ? selectedRide.pickup : defaultCenter;

  const path = selectedRide
    ? [selectedRide.pickup, selectedRide.dropoff]
    : [];

  return (
    <GoogleMap
      mapContainerStyle={mapContainerStyle}
      center={center}          
      zoom={13}
      onLoad={onMapLoad}
      options={{ disableDefaultUI: true }}
    >
      {rides.map(r => (
        <Marker
          key={`${r.id}-pickup`}
          position={r.pickup}
          icon="http://maps.google.com/mapfiles/ms/icons/green-dot.png"
        />
      ))}

      {rides.map(r => (
        <Marker
          key={`${r.id}-dropoff`}
          position={r.dropoff}
          icon="http://maps.google.com/mapfiles/ms/icons/red-dot.png"
        />
      ))}

      {path.length === 2 && (
        <Polyline
          path={path}
          options={{
            strokeColor: '#3b82f6',
            strokeOpacity: 0.8,
            strokeWeight: 4,
            clickable: true,           
          }}
          onClick={onPolylineClick}     
        />
      )}
    </GoogleMap>
  );
}
