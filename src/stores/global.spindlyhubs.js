import { ConnectHub } from "spindly-hubs";

const hub_name = 'GlobalHub';

export function GlobalHub(hub_instance_id, preserve = false) {
    let SpindlyStore = ConnectHub(hub_name, hub_instance_id, preserve);
    let SpindlyEventStore = (storename) => {
        let es = SpindlyStore(storename);
        return () => { es.set(Math.random()); };
    };
    return {
        AppName: SpindlyStore("AppName"),
        SaidHello: SpindlyEventStore("SaidHello"),
        HelloMessage: SpindlyStore("HelloMessage"),
        Events: SpindlyStore("Events"),
    }
}

export let Global = GlobalHub("Global", true);
