"use client"

import { useEffect, useState } from 'react';
import * as Packet from '../lib/packet'

export default function Home() {
    const [packetData, setPacketData] = useState<string | null>('test');

    console.log("Hey!1");

    useEffect(() => {
        console.log("Hey!2");
        const fetchData = async () => {
            console.log("Hey!3");
            try {
                const packet = Packet.NewPacket(11, 200, 0, "Test");
                const data: Uint8Array = await Packet.SendPacket(packet);
                const recv: Packet.Packet = Packet.FromBytes(data);
                setPacketData(recv.data);
            }
            catch (error) {
                console.error(error);
            }
        };

        fetchData();
    });

    return (
        <div>
            {packetData ? packetData : "No Data"}
        </div>
    );
}
