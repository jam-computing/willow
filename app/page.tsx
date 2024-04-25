import * as net from 'net';
import { Buffer } from 'buffer';

export default function Home() {
    return (
        <div>
            Hello, World!
        </div>
    );
}

function SendInitRequest() {

    const version = 1;
    const command = 1; // Init

    const vals = [
        200, 0, 0
    ];

    const nums = [
        Buffer.alloc(2),
        Buffer.alloc(2),
        Buffer.alloc(2)]
        .map((x) => {
            return new DataView(
                x.buffer,
                x.byteOffset,
                x.byteLength);
        });

    const zip: Uint16Array[] = (d: DataView[], n: number[]) => d.map((v, i) => v.setUint16(0, n[i], true));




    // const arr: Uint8Array = new Uint8Array(data);



    const client = new net.Socket();
    client.connect(8000, 'localhost', () => {
        console.log("Connected to server");
        client.write(arr);
    });
}
