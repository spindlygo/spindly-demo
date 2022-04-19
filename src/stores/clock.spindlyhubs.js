import { ConnectHub } from "spindly-hubs";

const hub_name = 'ClockHub';

export function ClockHub(hub_instance_id, preserve = false) {
    let SpindlyStore = ConnectHub(hub_name, hub_instance_id, preserve);
    let SpindlyEventStore = (storename) => {
        let es = SpindlyStore(storename);
        return () => { es.set(Math.random()); };
    };
    return {
        ClockFace: SpindlyStore("ClockFace"),
        TimeZone: SpindlyStore("TimeZone"),
    }
}

