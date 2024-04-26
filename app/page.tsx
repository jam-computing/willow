import * as Packet from '../lib/packet'
import '../app/card.css'

export default async function Home() {
    async function get(): Promise<string | null | undefined> {
        const fetchData = async () => {
            try {
                const packet = Packet.NewPacket(11, 200, 0, "Test");
                const data: Uint8Array = await Packet.SendPacket(packet);
                const recv: Packet.Packet = Packet.FromBytes(data);
                return recv.data;
            }
            catch (error) {
                console.error(error);
            }
        };
        return await fetchData();
    };

    const data = await get();

    const json = JSON.parse(data + "");

    return (
        <div className="p-4">
            <Card name={json.name} description={json.artist} />
        </div>
    );
}

interface CardProps {
    name: string,
    description: string
}

const Card: React.FC<CardProps> = ({ name, description }) => {
    return (
        <div className="playlist-card">
            <div className="name">{name}</div>
            <p className="description">{description}</p>
        </div>
    );
};
