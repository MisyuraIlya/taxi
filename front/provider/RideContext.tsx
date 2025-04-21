"use client";

import { createContext, useContext, ReactNode, useEffect, useRef } from 'react'
import { Ride } from '@/types/ride'
import { useRideStore } from '@/store/ride.store';
import { useNoitifcation } from './NotificationProvider';

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
}

const RideContext = createContext<RideContextValue | undefined>(undefined)

export function RideProvider({ children }: { children: ReactNode }) {
  const role = useRideStore((state) => state.role)
  const userId = useRideStore((state) => state.userId)
  const rides = useRideStore((state) => state.rides)
  const selectedRideId = useRideStore((state) => state.selectedRideId)
  const pendingRide = useRideStore((state) => state.pendingRide)
  const events = useRideStore((state) => state.events)
  const { message } = useNoitifcation()
  const wsRef = useRef<WebSocket | null>(null)

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
        message('message',e.data)
        const msg = JSON.parse(e.data)
        console.log('msg',msg)
        if (role === 'driver' && msg.type === 'rideRequest') {
          useRideStore.getState()._setPendingRide(msg.ride)
        } else if (role === 'client') {
          const entry = msg.type + (msg.data ? `: ${JSON.stringify(msg.data)}` : '')
          useRideStore.getState()._addEvent(entry)
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
    if (!socket || !ride) return
    socket.send(JSON.stringify({ type: 'acceptRide', rideId: ride.id }))
    useRideStore.getState()._setPendingRide(null)
  }

  const declineRide = () => {
    const socket = wsRef.current
    const ride = pendingRide
    if (!socket || !ride) return
    socket.send(JSON.stringify({ type: 'declineRide', rideId: ride.id }))
    useRideStore.getState()._setPendingRide(null)
  }

  const requestMatch = () => {
    wsRef.current?.send(JSON.stringify({ type: 'requestRide' }))
  }

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
  }

  return (
    <RideContext.Provider value={contextValue}>
      {children}
    </RideContext.Provider>
  )
}

export function useRide() {
  const ctx = useContext(RideContext)
  if (!ctx) throw new Error('useRide must be used within RideProvider')
  return ctx
}
