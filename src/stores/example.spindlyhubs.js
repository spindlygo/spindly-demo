import { ConnectHub } from "spindly-hubs";

const hub_name = 'ExampleHub';

export function ExampleHub(hub_instance_id, preserve = false) {
    let SpindlyStore = ConnectHub(hub_name, hub_instance_id, preserve);
    let SpindlyEventStore = (storename) => {
        let es = SpindlyStore(storename);
        return () => { es.set(Math.random()); };
    };
    return {
        Name: SpindlyStore("Name"),
        Greating: SpindlyStore("Greating"),
    }
}

