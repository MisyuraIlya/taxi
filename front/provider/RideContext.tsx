"use client";

import { createContext, useContext, useState, useEffect, ReactNode } from "react";
import { useSearchParams } from "next/navigation";
import { Ride } from "@/types/Ride";

interface RideContextValue {
  rides: Ride[];
  selectedRideId: string | null;
  setSelectedRideId: (id: string | null) => void;
  pendingRide: Ride | null;
  events: string[];
  acceptRide: () => void;
  declineRide: () => void;
  requestMatch: () => void;
}

const RideContext = createContext<RideContextValue | undefined>(undefined);

export function RideProvider({ children }: { children: ReactNode }) {
  const params = useSearchParams();
  const roleParam = params.get("role");
  const userId = params.get("userId") || "";
  const role = roleParam === "driver" || roleParam === "client" ? roleParam : null;

  const [rides] = useState<Ride[]>([
    { id: "ride1", pickup: { lat: 37.779, lng: -122.418 }, dropoff: { lat: 37.784, lng: -122.409 }, timestamp: "2025-04-18T10:30:00Z" },
    { id: "ride2", pickup: { lat: 37.772, lng: -122.421 }, dropoff: { lat: 37.768, lng: -122.415 }, timestamp: "2025-04-17T15:45:00Z" },
  ]);
  const [selectedRideId, setSelectedRideId] = useState<string | null>(null);
  const [pendingRide, setPendingRide] = useState<Ride | null>(null);
  const [events, setEvents] = useState<string[]>([]);
  const [ws, setWs] = useState<WebSocket | null>(null);

  useEffect(() => {
    if (!role || !userId) return;
    const socket = new WebSocket("ws://localhost:4000");
    socket.onopen = () => socket.send(JSON.stringify({ type: "init", role, userId }));
    socket.onmessage = (e) => {
      const msg = JSON.parse(e.data);
      if (role === "driver" && msg.type === "rideRequest") setPendingRide(msg.ride);
      if (role === "client") setEvents((prev) => [...prev, msg.type + (msg.data ? `: ${JSON.stringify(msg.data)}` : "")]);
    };
    setWs(socket);
    return () => socket.close();
  }, [role, userId]);

  const acceptRide = () => {
    if (!ws || !pendingRide) return;
    ws.send(JSON.stringify({ type: "acceptRide", rideId: pendingRide.id }));
    setPendingRide(null);
  };

  const declineRide = () => {
    if (!ws || !pendingRide) return;
    ws.send(JSON.stringify({ type: "declineRide", rideId: pendingRide.id }));
    setPendingRide(null);
  };

  const requestMatch = () => {
    ws?.send(JSON.stringify({ type: "requestRide" }));
  };

  return (
    <RideContext.Provider value={{ rides, selectedRideId, setSelectedRideId, pendingRide, events, acceptRide, declineRide, requestMatch }}>
      {children}
    </RideContext.Provider>
  );
}

export function useRide() {
  const ctx = useContext(RideContext);
  if (!ctx) throw new Error("useRide must be used within RideProvider");
  return ctx;
}
