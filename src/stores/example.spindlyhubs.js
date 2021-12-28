import { ConnectHub } from "spindly-hubs";

const hub_name = 'ExampleHub';

export function ExampleHub(hub_instance_id) {
    let SpindlyStore = ConnectHub(hub_name, hub_instance_id);
    let SpindlyEventStore = (storename) => {
        let es = SpindlyStore(storename);
        return () => { es.set(Math.random()); };
    };
    return {
        Message: SpindlyStore("Message"),
        Greating: SpindlyStore("Greating"),
    }
}

