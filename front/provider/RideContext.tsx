"use client";

import { createContext, useContext, ReactNode, useEffect, useRef, useState } from 'react'
import { Ride } from '@/types/ride'
import { useRideStore } from '@/store/ride.store';
import { useNoitifcation } from './NotificationProvider';
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Button } from '@/components/ui/button';
import { status } from '@grpc/grpc-js';

interface RideContextValue {
  role: 'driver' | 'client' | null
  setRole: (role: 'driver' | 'client' | null) => void
  userId: string | null
  setUserId: (id: string | null) => void
  rides: Ride[]
  selectedRideId: string | null
  setSelectedRideId: (id: string | null) => void
  pendingRide: Ride | null
  events: string[]
  acceptRide: () => void
  declineRide: () => void
  requestMatch: () => void
  matchService: () => void
  updateLocation: (userId: string, latitude: string | number , longitude: string | number) => void
  testDriver: () => void
}

const RideContext = createContext<RideContextValue | undefined>(undefined)

export function RideProvider({ children }: { children: ReactNode }) {
  const [dialogOpen, setDialogOpen] = useState(false) 
  const [isInRide, setIsInRide] = useState(false)
  const role = useRideStore((state) => state.role)
  const userId = useRideStore((state) => state.userId)
  const rides = useRideStore((state) => state.rides)
  const selectedRideId = useRideStore((state) => state.selectedRideId)
  const pendingRide = useRideStore((state) => state.pendingRide)
  const events = useRideStore((state) => state.events)
  const { message } = useNoitifcation()
  const wsRef = useRef<WebSocket | null>(null)
  const { setDrivers, latitude, longitude, drivers, driverId, setDriverLatitude, setDriverLongitude, setDriverId } = useRideStore()


  useEffect(() => {
    if (!role || !userId) return

    const base = process.env.NEXT_PUBLIC_WS_URL || 'ws://localhost:8082/ws'
    const paramKey = role === 'driver' ? 'driver_id' : 'client_id'
    const socketUrl = `${base}?${paramKey}=${encodeURIComponent(userId)}`

    const socket = new WebSocket(socketUrl)
    wsRef.current = socket

    socket.onopen = () => {
      socket.send(JSON.stringify({ type: 'init', role, userId }))
    }

    socket.onmessage = (e) => {
      try {
        console.log('Received message:', e.data)
        
        const msg = JSON.parse(e.data)
        console.log('msg',msg)
        if (role === 'driver' && msg.message.type === 'rideRequest') {
          message(msg.message.type,msg.message.data)
          setDialogOpen(true)
        } else if (role === 'client') {
          message(msg.message.type,msg.message.data)
        }
      } catch {
        useRideStore.getState()._setPendingRide(JSON.parse(e.data) as Ride)
      }
    }

    return () => {
      socket.close()
      wsRef.current = null
    }
  }, [role, userId])

  const acceptRide = () => {
    const socket = wsRef.current
    const ride = pendingRide
    if (!socket) return
    socket.send('accept')
  }

  const declineRide = () => {
    const socket = wsRef.current
    const ride = pendingRide
    if (!socket || !ride) return
    socket.send(JSON.stringify({ type: 'declineRide', rideId: ride.id }))
  }

  const handleDriverLocation =  async () => {
    if (!userId || !role) return
    try {
      const res = await fetch(
        `/api/driver/${driverId}`
      )
      if (!res.ok) throw new Error(await res.text())
      const data = await res.json()
      console.log('driver location',data)
      setIsInRide(true)
      setDriverLatitude(+data.latitude)
      setDriverLongitude(+data.longitude)
    } catch (e) {
      console.error('requestMatch error', e)
    }
  } 

  const requestMatch = async () => {
    if (!userId || !role) return
  
    try {
      const res = await fetch(
        `/api/drivers?latitude=37.7749` +
        `&longitude=-122.4194` +
        `&radius=1000&limit=20&status=active`
      )
      if (!res.ok) throw new Error(await res.text())
      const data = await res.json()
      setDrivers(data.drivers)

    } catch (e) {
      console.error('requestMatch error', e)
      useRideStore.getState()._addEvent(`Match error: ${(e as Error).message}`)
    }
  }

  const matchService = async () => {
    if (!userId || !role) return
    try {
      const res = await fetch(
        `/api/matchService`
      ,{
        method:'POST',
        body:JSON.stringify({longitude, latitude, radius:100000, limit:10})
      })
      if (!res.ok) {
        message("driver on the way","no drivers found around")
      }
      const data = await res.json()
      message("driver on the way",data?.drivers[0]?.driverId)
      const deleteDriver = drivers.filter((x) => x.driverId !== data?.drivers[0]?.driverId)
      setDriverId(data?.drivers[0]?.driverId)
      setDriverLatitude(data?.drivers[0]?.latitude)
      setDriverLongitude(data?.drivers[0]?.longitude)
      setDrivers(deleteDriver)
    } catch (e) {
      console.error('requestMatch error', e)
      message("error in match",``)
    }
  }

  const updateLocation = async (userId: string,latitude: string | number , longitude: string | number) => {
    console.log('updateLocation start')
    let obj = {
      latitude: latitude,
      longitude: longitude,
      driverId: userId,
      status: 'active'
    }

    try {
      const res = await fetch(
        `/api/driver/${userId}`, {
          method: 'PUT',
          body: JSON.stringify(obj)
        }
      )
      if (!res.ok) throw new Error(await res.text())
      const data = await res.json()
      console.log('updateLocation',data)
    } catch (e) {
      console.error('requestMatch error', e)
    }
  }

  const testDriver = async () => {
    // 1. Configuration
    const driverId       = 'driver123'
    const driverStart    = { latitude:  37.773487, longitude: -122.418687 }
    const clientLoc      = { latitude:  37.769183, longitude: -122.407994 }
    const numRoutePoints = 20      
    const intervalMs     = 5000    
  
    // 2. Build an array of intermediate coords
    function generateRoute(
      start: { latitude: number, longitude: number },
      end:   { latitude: number, longitude: number },
      steps: number
    ) {
      const route = []
      const latStep = (end.latitude  - start.latitude)  / steps
      const lngStep = (end.longitude - start.longitude) / steps
  
      for (let i = 0; i <= steps; i++) {
        route.push({
          latitude:  +(start.latitude  + latStep * i).toFixed(6),
          longitude: +(start.longitude + lngStep  * i).toFixed(6),
        })
      }
  
      return route
    }
  
    // 3. Start the “movement” loop
    const mockRoute = generateRoute(driverStart, clientLoc, numRoutePoints)
    let idx = 0
  
    setInterval(() => {
      const { latitude, longitude } = mockRoute[idx]
      updateLocation(driverId, latitude, longitude)
      idx = (idx + 1) % mockRoute.length
    }, intervalMs)
  }


  useEffect(() => {
    requestMatch()
  },[userId,role])

  useEffect(() => {
    if (role === 'client' && userId) {
      handleDriverLocation();
      const intervalId = setInterval(() => {
        handleDriverLocation();
      }, 5000);
      return () => clearInterval(intervalId);
    } else {
      setIsInRide(false);
    }
  }, [driverId]);
  

  const contextValue: RideContextValue = {
    role,
    setRole: (role) => useRideStore.getState().setRole(role),
    userId,
    setUserId: (id) => useRideStore.getState().setUserId(id),
    rides,
    selectedRideId,
    setSelectedRideId: (id) => useRideStore.getState().setSelectedRideId(id),
    pendingRide,
    events,
    acceptRide,
    declineRide,
    requestMatch,
    matchService,
    updateLocation,
    testDriver
  }

  return (
    <RideContext.Provider value={contextValue}>

      <Dialog open={dialogOpen} onOpenChange={setDialogOpen}>
        <DialogContent className="sm:max-w-md">
          <DialogHeader>
            <DialogTitle>New Ride Request</DialogTitle>
            <DialogDescription>
              A client is requesting a ride. Would you like to accept?
            </DialogDescription>
          </DialogHeader>
          <DialogFooter className="sm:justify-end">
            <Button onClick={() => {
              acceptRide()
              setDialogOpen(false)
            }}>
              Accept
            </Button>
            <Button variant="secondary" onClick={() => {
              declineRide()
              setDialogOpen(false)
            }}>
              Decline
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
      
      {children}
    </RideContext.Provider>
  )
}

export function useRide() {
  const ctx = useContext(RideContext)
  if (!ctx) throw new Error('useRide must be used within RideProvider')
  return ctx
}
