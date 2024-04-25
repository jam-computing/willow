import * as Packet from '../lib/packet'

export default function Home() {
    const packet = Packet.NewPacket(11, 200, 0, "Test");

    (async () => {
        try {
            const data: Uint8Array = await Packet.SendPacket(packet);
            console.log(data.length);
            data.forEach((x, i) => {
                console.log(`Byte ${i}: ${x}`);
            });
            const recv: Packet.Packet = Packet.FromBytes(data);
            console.log("Received Version:", recv.version);
            console.log("Received Command:", recv.command);
            console.log("Received Status:", recv.status);
            console.log("Received Id:", recv.id);
            console.log("Received Data:", recv.data);
        }
        catch (error) {
            console.error('error receiving error: ', error);
        }
    })();

    return (
        <div>
        </div>
    );
}
