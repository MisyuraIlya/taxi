import { create } from 'zustand'
import { persist, createJSONStorage, PersistOptions } from 'zustand/middleware'
import { Ride } from '@/types/ride'

// New Driver interface
export interface Driver {
  driverId: string
  latitude: number
  longitude: number
}

interface RideState {
  role: 'driver' | 'client' | null
  userId: string | null
  rides: Ride[]
  selectedRideId: string | null
  pendingRide: Ride | null
  events: string[]

  // Driver-related state
  drivers: Driver[]
  selectedDriverId: string | null

  // Global location state
  latitude: number | null
  longitude: number | null
  driverLatitude: number | null
  driverLongitude: number | null

  setDriverLongitude: (val: number) => void
  setDriverLatitude: (val: number) => void

  

  setRole: (role: 'driver' | 'client' | null) => void
  setUserId: (id: string | null) => void
  setSelectedRideId: (id: string | null) => void
  _setPendingRide: (ride: Ride | null) => void
  _addEvent: (entry: string) => void

  // Driver setters
  setDrivers: (drivers: Driver[]) => void
  addDriver: (driver: Driver) => void
  updateDriver: (driver: Driver) => void
  removeDriver: (driverId: string) => void
  setSelectedDriverId: (driverId: string | null) => void

  // Location setters
  setLatitude: (latitude: number | null) => void
  setLongitude: (longitude: number | null) => void
  setLocation: (latitude: number | null, longitude: number | null) => void
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

      drivers: [],
      selectedDriverId: null,

      // Initialize location state
      latitude: null,
      longitude: null,

      //
      driverLatitude: null,
      driverLongitude: null,

      setDriverLatitude: (driverLatitude) => set({driverLatitude}),
      setDriverLongitude: (driverLongitude) => set({driverLongitude}),

      setRole: (role) => set({ role }),
      setUserId: (userId) => set({ userId }),
      setSelectedRideId: (id) => set({ selectedRideId: id }),

      _setPendingRide: (ride) => set({ pendingRide: ride }),
      _addEvent: (entry) => set((state) => ({ events: [...state.events, entry] })),

      setDrivers: (drivers) => set({ drivers }),
      addDriver: (driver) => set((state) => ({ drivers: [...state.drivers, driver] })),
      updateDriver: (driver) => set((state) => ({
        drivers: state.drivers.map((d) => d.driverId === driver.driverId ? driver : d)
      })),
      removeDriver: (driverId) => set((state) => ({
        drivers: state.drivers.filter((d) => d.driverId !== driverId)
      })),
      setSelectedDriverId: (driverId) => set({ selectedDriverId: driverId }),

      // Location setters
      setLatitude: (latitude) => set({ latitude }),
      setLongitude: (longitude) => set({ longitude }),
      setLocation: (latitude, longitude) => set({ latitude, longitude }),
    }),
    {
      name: 'ride-storage',
      storage: createJSONStorage(() => localStorage),
    } as PersistOptions<RideState, RideState>
  )
)
