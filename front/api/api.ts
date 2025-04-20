import axios from 'axios';

const gateway = process.env.GATEWAY_URL;

export async function findDrivers(lat: number, lng: number, radius = 1000, limit = 10) {
  const resp = await axios.post(`${gateway}/findDrivers`, { latitude: lat, longitude: lng, radius, limit });
  return resp.data.drivers;
}

export async function createOrder(payload: any) {
  const resp = await axios.post(`${gateway}/createOrder`, payload);
  return resp.data;
}
