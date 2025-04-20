"use client";

import { useEffect, useState } from "react";
import RideMap from "@/components/RideMap";
import { Ride } from "@/types/Ride";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogFooter,
  DialogTitle,
} from "@/components/ui/dialog";

interface MapPageProps {
  role: "driver" | "client";
  userId: string;
}

export default function MapPage({ role, userId }: MapPageProps) {
  const [rides] = useState<Ride[]>([
    {
      id: "ride1",
      pickup: { lat: 37.779, lng: -122.418 },
      dropoff: { lat: 37.784, lng: -122.409 },
      timestamp: "2025-04-18T10:30:00Z",
    },
    {
      id: "ride2",
      pickup: { lat: 37.772, lng: -122.421 },
      dropoff: { lat: 37.768, lng: -122.415 },
      timestamp: "2025-04-17T15:45:00Z",
    },
  ]);

  const [selectedRideId, setSelectedRideId] = useState<string | null>(null);
  const [ws, setWs] = useState<WebSocket | null>(null);
  const [pendingRide, setPendingRide] = useState<Ride | null>(null);
  const [events, setEvents] = useState<string[]>([]);

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:4000");
    socket.onopen = () => {
      socket.send(JSON.stringify({ type: "init", role, userId }));
    };
    socket.onmessage = (e) => {
      const msg = JSON.parse(e.data);
      if (role === "driver" && msg.type === "rideRequest") {
        setPendingRide(msg.ride);
      }
      if (role === "client") {
        setEvents((prev) => [...prev, msg.type + (msg.data ? `: ${JSON.stringify(msg.data)}` : "")]);
      }
    };
    setWs(socket);
    return () => socket.close();
  }, [role, userId]);

  const acceptRide = () => {
    if (pendingRide) {
      ws?.send(JSON.stringify({ type: "acceptRide", rideId: pendingRide.id }));
      setPendingRide(null);
    }
  };
  const declineRide = () => {
    if (pendingRide) {
      ws?.send(JSON.stringify({ type: "declineRide", rideId: pendingRide.id }));
      setPendingRide(null);
    }
  };

  const requestMatch = () => {
    ws?.send(JSON.stringify({ type: "requestRide" }));
  };

  return (
    <div className="flex flex-col h-screen">
      <div className="flex-1">
        <RideMap
          rides={rides}
          selectedRideId={selectedRideId}
          onPolylineClick={() => setSelectedRideId(null)}
        />
      </div>
      <div className="p-4 bg-gray-50 dark:bg-gray-800 space-y-4 overflow-auto">
        {role === "driver" && pendingRide && (
          <Dialog open={!!pendingRide} onOpenChange={() => setPendingRide(null)}>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>New Ride Request</DialogTitle>
              </DialogHeader>
              <p>
                Ride from ({pendingRide.pickup.lat.toFixed(4)}, {pendingRide.pickup.lng.toFixed(4)}) to ({pendingRide.dropoff.lat.toFixed(4)}, {pendingRide.dropoff.lng.toFixed(4)})
              </p>
              <DialogFooter className="space-x-2">
                <Button onClick={acceptRide}>Accept</Button>
                <Button variant="outline" onClick={declineRide}>Decline</Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        )}

        {role === "client" && (
          <Button onClick={requestMatch} className="w-full">
            Match Ride
          </Button>
        )}

        {role === "client" && events.length > 0 && (
          <div className="space-y-2">
            <h4 className="text-lg font-medium">Events:</h4>
            {events.map((evt, idx) => (
              <div key={idx} className="text-sm">
                {evt}
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
