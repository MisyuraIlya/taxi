export const runtime = 'nodejs'

import { NextResponse } from 'next/server'
import { geoClient }     from '@/lib/grpcClient'
import grpc, { ServiceError } from '@grpc/grpc-js'

export async function GET(request: Request) {
  const { searchParams } = new URL(request.url)

  const latitudeRaw  = searchParams.get('latitude')
  const longitudeRaw = searchParams.get('longitude')
  const radiusRaw    = searchParams.get('radius') ?? '5'
  const limitRaw     = searchParams.get('limit')  ?? '10'
  const status       = searchParams.get('status') ?? 'active'

  if (!latitudeRaw || !longitudeRaw) {
    return NextResponse.json(
      { error: 'latitude and longitude are required' },
      { status: 400 }
    )
  }

  const payload = {
    latitude:  parseFloat(latitudeRaw),
    longitude: parseFloat(longitudeRaw),
    radius:    parseFloat(radiusRaw),
    limit:     parseInt(limitRaw, 10),
    status,
  }

  return new Promise<NextResponse>((resolve) => {
    geoClient.FindDrivers(payload, (err: grpc.ServiceError | null, reply: any) => {
      if (err) {
        console.error('[gRPC FindDrivers]', err)
        resolve(
          NextResponse.json({ error: err.message }, { status: 500 })
        )
      } else {
        resolve(NextResponse.json(reply))
      }
    })
  })
}
