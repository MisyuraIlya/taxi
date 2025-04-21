"use client";

import RideMap from "@/components/RideMap";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogFooter,
  DialogTitle,
} from "@/components/ui/dialog";
import { useRide } from "@/provider/RideContext";

export default function MapPage() {
  const {
    role,
    rides,
    selectedRideId,
    setSelectedRideId,
    pendingRide,
    events,
    acceptRide,
    declineRide,
    requestMatch,
  } = useRide();

  return (
    <div className="flex flex-col h-screen">
      <div className="flex-1">
        <RideMap
          rides={rides}
          selectedRideId={selectedRideId}
          onPolylineClick={() => setSelectedRideId(null)}
        />
      </div>
      <div className="p-12 bg-gray-50 dark:bg-gray-800 space-y-4 overflow-auto">
        {role === "driver" && pendingRide && (
          <Dialog open>
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
