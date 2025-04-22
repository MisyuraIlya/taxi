import grpc from '@grpc/grpc-js'
import protoLoader from '@grpc/proto-loader'
import path from 'path'

const PROTO_PATH = path.resolve(process.cwd(), 'proto/geo.proto')
const packageDef = protoLoader.loadSync(PROTO_PATH, {/* …options… */})
const { GeoService } = (grpc.loadPackageDefinition(packageDef) as any).geo

// Read from env, fallback to 127.0.0.1:50051 if missing
const addr = process.env.GEO_SERVICE_ADDR ?? 'localhost:8081'

export const geoClient = new GeoService(
  addr,
  grpc.credentials.createInsecure()
)