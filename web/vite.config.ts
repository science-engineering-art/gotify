import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: '0.0.0.0',
    port: 80,
  }
})

import os from 'os';
import * as dgram from 'dgram'
import * as net from 'net';

function numberToBytes(number: number): Uint8Array {
  // const len = Math.ceil(Math.log2(number) / 8);
  const byteArray = new Uint8Array(1);

  for (let index = 0; index < byteArray.length; index++) {
      const byte = number & 0xff;
      byteArray[index] = byte;
      number = (number - byte) / 256;
  }

  return byteArray;
}

const socket = dgram.createSocket("udp4")

const ip = Object.values(os.networkInterfaces())
  .flat()
  .filter(net => net?.family === 'IPv4' && !net.internal)
  .pop()?.address

const parts = ip?.split(".");

const IPv4 = parts?.map(Number)?.map(numberToBytes)

const resp = new Uint8Array(4)
resp.set(IPv4?.[0]!)
resp.set(IPv4?.[1]!, 1)
resp.set(IPv4?.[2]!, 2)
resp.set(IPv4?.[3]!, 3)

socket.on("message", (_, rinfo) => {
  const tcpSocket = {
    port: rinfo.port,
    host: rinfo.address,
  }
  const client = net.createConnection(tcpSocket, () => {
    console.log('Connected to DNS');
  })

  client.on('connect', ()=>{ 
    client.write(resp); 
  })

  client.on('error', (err) => {
    console.log(err)
  })
});

socket.bind(41234);
