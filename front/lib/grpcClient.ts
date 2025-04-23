// lib/grpcClient.ts
import grpc from '@grpc/grpc-js'
import protoLoader from '@grpc/proto-loader'
import path from 'path'

const PROTO_DIR        = path.resolve(process.cwd(), 'proto')
const GEO_PROTO_PATH   = path.join(PROTO_DIR, 'geo.proto')
const MATCH_PROTO_PATH = path.join(PROTO_DIR, 'matching.proto')

const loaderOptions = {
  keepCase: false,
  longs:   String,
  enums:   String,
  defaults: true,
  oneofs:  true,
}

const geoDef   = protoLoader.loadSync(GEO_PROTO_PATH, loaderOptions)
const matchDef = protoLoader.loadSync(MATCH_PROTO_PATH, loaderOptions)

const grpcGeo   = grpc.loadPackageDefinition(geoDef)   as any
const grpcMatch = grpc.loadPackageDefinition(matchDef) as any

// From geo.proto:
//    package geo;
//    service GeoService { … }
const { GeoService } = grpcGeo.geo

// From matching.proto:
//    package matching;
//    service MatchingService { … }
const { MatchingService } = grpcMatch.matching

export const geoClient = new GeoService(
  process.env.GEO_SERVICE_ADDR   ?? 'localhost:8081',
  grpc.credentials.createInsecure()
)

export const matchClient = new MatchingService(
  process.env.MATCH_SERVICE_ADDR ?? 'localhost:9090',
  grpc.credentials.createInsecure()
)
