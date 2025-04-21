import { create } from 'zustand'
import { persist, createJSONStorage, PersistOptions } from 'zustand/middleware'
import { Ride } from '@/types/ride'

interface RideState {
  role: 'driver' | 'client' | null
  userId: string | null
  rides: Ride[]
  selectedRideId: string | null
  pendingRide: Ride | null
  events: string[]
  setRole: (role: 'driver' | 'client' | null) => void
  setUserId: (id: string | null) => void
  setSelectedRideId: (id: string | null) => void
  _setPendingRide: (ride: Ride | null) => void
  _addEvent: (entry: string) => void
}

export const useRideStore = create<RideState>()(
  persist(
    (set) => ({
      role: null,
      userId: null,
      rides: [
        { id: 'ride1', pickup: { lat: 37.779, lng: -122.418 }, dropoff: { lat: 37.784, lng: -122.409 }, timestamp: '2025-04-18T10:30:00Z' },
        { id: 'ride2', pickup: { lat: 37.772, lng: -122.421 }, dropoff: { lat: 37.768, lng: -122.415 }, timestamp: '2025-04-17T15:45:00Z' }
      ],
      selectedRideId: null,
      pendingRide: null,
      events: [],

      setRole: (role) => set({ role }),
      setUserId: (userId) => set({ userId }),
      setSelectedRideId: (id) => set({ selectedRideId: id }),

      // internal setters used by context
      _setPendingRide: (ride) => set({ pendingRide: ride }),
      _addEvent: (entry) => set((state) => ({ events: [...state.events, entry] })),
    }),
    {
      name: 'ride-storage',
      storage: createJSONStorage(() => localStorage),
    } as PersistOptions<RideState, RideState>
  )
)
