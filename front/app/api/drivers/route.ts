import type { NextApiRequest, NextApiResponse } from 'next'
import { geoClient } from '@/lib/grpcClient'
import { NextResponse } from 'next/server'

export async function GET(req: NextApiRequest, res: NextApiResponse) {
    const { latitude, longitude, radius = '5', limit = '10', status = 'active' } = req.query
    const payload = {
      latitude:  parseFloat(latitude as string),
      longitude: parseFloat(longitude as string),
      radius:    parseFloat(radius as string),
      limit:     parseInt(limit as string, 10),
      status:    status as string,
    }
  
    geoClient.FindDrivers(payload, (err: Error|null, reply: any) => {
      if (err) {
        console.error('[gRPC FindDrivers]', err)
        return res.status(500).json({ error: err.message })
      }
      // reply.drivers is your array of geo.Driver
      res.status(200).json(reply)
    })
}