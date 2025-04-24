// app/api/driver/[driverId]/route.ts
import { NextRequest, NextResponse } from 'next/server'
import { geoClient } from '@/lib/grpcClient'

export async function GET(
  request: NextRequest,
  context: { params: Promise<{ driverId: string }> }
) {
  const { driverId } = await context.params

  try {
    const reply = await new Promise<any>((resolve, reject) => {
      geoClient.GetLocation({ driverId }, (err:any, data:any) => {
        if (err) {
          reject(err)
        } else {
          resolve(data)
        }
      })
    })

    return NextResponse.json(reply)
  } catch (err: any) {
    console.error('[gRPC GetLocation]', err)
    return NextResponse.json({ error: err.message }, { status: 500 })
  }
}
