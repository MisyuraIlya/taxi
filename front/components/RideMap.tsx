"use client";
import { useLoadScript, GoogleMap, Marker, Polyline } from "@react-google-maps/api";
import { useRef, useCallback } from "react";
import { Ride } from "@/types/ride";
import { useRideStore } from "@/store/ride.store";

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
  const { drivers, latitude, longitude} = useRideStore();
  const { isLoaded, loadError } = useLoadScript({
    googleMapsApiKey: process.env.NEXT_PUBLIC_GMAPS_KEY!,
  });

  const mapRef = useRef<google.maps.Map | null>(null);
  const onMapLoad = useCallback((map: google.maps.Map) => {
    mapRef.current = map;
  }, []);

  if (loadError) return <div>Error loading map</div>;
  if (!isLoaded)  return <div>Loading map…</div>;

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
      {/* Ride pickup/dropoff pins */}
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

      {/* Driver locations */}
      {drivers.map(d => (
        <Marker
          key={d.driverId}
          position={{ lat: d.latitude, lng: d.longitude }}
          icon="http://maps.google.com/mapfiles/ms/icons/blue-dot.png"
        />
      ))}

      <Marker
        position={{ lat: latitude!, lng: longitude! }}
        icon={{
          url: "http://maps.google.com/mapfiles/kml/shapes/man.png",
          scaledSize: new window.google.maps.Size(32, 32),
          anchor: new window.google.maps.Point(16, 32),
        }}
      />

      {/* Selected-ride path */}
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
