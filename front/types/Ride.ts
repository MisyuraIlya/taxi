export interface Ride {
    id: string;
    pickup: { lat: number; lng: number };
    dropoff: { lat: number; lng: number };
    timestamp: string; // ISO or formatted date
  }
  