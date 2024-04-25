import * as net from 'net';

const PROTOCOL_VERSION: number = 1;

export interface Packet {
    version: number;
    command: number;
    status: number;
    id: number;
    data: string | null;
}

export function SendPacket(packet: Packet): Promise<Uint8Array> {
    const vals = [
        packet.version,
        packet.command,
        packet.status,
        packet.id,
        0
    ];

    const bytes = [];

    for(let i = 0; i < vals.length; i++) {
        if(i < 2) {
            bytes.push(vals[i]);
        }
        else {
            const high = (vals[i] >> 8) & 0xFF;
            const low = vals[i] & 0xFF;
            bytes.push(low, high);
        }
    }

    const client = new net.Socket();
    const buf = new Uint8Array(bytes);

    return new Promise((resolve, reject) => {
        client.connect(8080, 'localhost', () => {
            console.log("Connected to server");
            client.write(buf);
        });

        client.on('data', (data) => {
            const arr = new Uint8Array(data);
            resolve(arr);
        });

        client.on('error', (error) => {
            reject(error);
        });
    });
}

export function NewPacket(command: number, status: number, id: number, data: string | null): Packet {
    const newPacket: Packet = {
        version: PROTOCOL_VERSION,
        command: command,
        status: status,
        id: id,
        data: data
    };

    return newPacket;
}

export function FromBytes(bytes: Uint8Array): Packet {
    const version = bytes[0];
    const command = bytes[0];
    const status = (bytes[2] << 8) | bytes[3];
    const id = (bytes[4] << 8) | bytes[5];
    const length = (bytes[6] << 8) | bytes[7];
    const data = length > 0 ? new TextDecoder().decode(bytes.slice(8)) : null;

    return {
        version,
        command,
        status,
        id,
        data
    };
}
