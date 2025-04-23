// app/api/match/route.ts
export const runtime = 'nodejs'

import { NextResponse } from 'next/server'
import { matchClient }  from '@/lib/grpcClient'

export async function POST(request: Request) {
  let body: any
  try {
    body = await request.json()
  } catch {
    return NextResponse.json({ error: 'Invalid JSON' }, { status: 400 })
  }

  const {
    latitude,
    longitude,
    radius = 5,
    limit  = 10,
  } = body

  if (latitude == null || longitude == null) {
    return NextResponse.json(
      { error: 'latitude and longitude are required' },
      { status: 400 }
    )
  }

  const payload = {
    latitude:  Number(latitude),
    longitude: Number(longitude),
    radius:    Number(radius),
    limit:     parseInt(limit, 10),
  }

  return new Promise<NextResponse>((resolve) => {
    matchClient.MatchClients(
      payload,
      (err:any, reply:any) => {
        if (err) {
          console.error('[gRPC MatchClients]', err)
          resolve(
            NextResponse.json({ error: err.message }, { status: 500 })
          )
        } else {
          resolve(NextResponse.json(reply))
        }
      }
    )
  })
}
