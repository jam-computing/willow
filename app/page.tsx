import * as Packet from '../lib/packet'

export default async function Home() {
    async function get(): Promise<string | null | undefined> {
        const fetchData = async () => {
            try {
                const packet = Packet.NewPacket(11, 200, 0, "Test");
                const data: Uint8Array = await Packet.SendPacket(packet);
                const recv: Packet.Packet = Packet.FromBytes(data);
                console.log(recv.data);
                return recv.data;
            }
            catch (error) {
                console.error(error);
            }
        };

        return await fetchData();
    };

    const data = await get();
    console.log(data);

    return (
        <div>
            {data}
        </div>
    );
}
