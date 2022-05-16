<template>
    <div class="flex h-full flex-col md:flex-row">
        <div class="flex h-full grow flex-col justify-between p-4 pb-12 md:p-10">
            <Navbar :status="status" />
            <Main :downloadStats="stats.down" :uploadStats="stats.up" />
            <Stats
                :downloadStats="stats.down"
                :uploadStats="stats.up"
                :latencyStats="stats.latency"
            />
        </div>

        <Sidebar :devices="stats.devices" :outages="stats.outages" />
    </div>
</template>

<script setup lang="ts">
const LOOKBACK = 120;

const status = ref({
    ws: "offline",
    dishy: "offline",
    uplink: "offline",
})
const config = ref({
    selection: "live",
});
const stats = ref({
    down: new Map<Date, number>(),
    up: new Map<Date, number>(),
    latency: new Map<Date, number>(),
    devices: [] as any[],
    outages: new Map<string, Map<Date, number>>(),
});

function addPoint(map: Map<Date, number>, timestamp: Date, value: number) {
    map.set(timestamp, value);
    if (map.size > LOOKBACK) map.delete(map.keys().next().value);
}

function connectWS() {
    let schema = "ws";
    if (window.location.protocol === "https:") {
        schema = "wss";
    }
    let ws = new WebSocket(schema + "://"+ window.location.host + "/ws/omnibus/live");
    ws.onopen = () => {
        status.value.ws = "online";
    };
    ws.onclose = () => {
        status.value.ws = "offline";
        setTimeout(function() {
            connectWS();
        }, 1000);
    };
    ws.onerror = () => {
        status.value.ws = "unknown";
        ws.close();
    };
    ws.onmessage = (event: any) => {
        var message = JSON.parse(event.data);
        if (message.error) {
            console.log("WS error", message.error);
            ws.close();
            return;
        }

        const data = message.data;
        switch(message.type) {
            case "INIT":
                const down = new Map<Date, number>();
                const up = new Map<Date, number>();
                const latency = new Map<Date, number>();
                if (data.down) {
                    data.down.forEach((point: any) => {
                        addPoint(down, new Date(point.timestamp), point.value / 1000000);
                    });
                }
                if (data.up) {
                    data.up.forEach((point: any) => {
                        addPoint(up, new Date(point.timestamp), point.value / 1000000);
                    });
                }
                if (data.latency) {
                    data.latency.forEach((point: any) => {
                        addPoint(latency, new Date(point.timestamp), point.value);
                    });
                }
                stats.value.down = down;
                stats.value.up = up;
                stats.value.latency = latency;
                break;
            case "CHART_UPDATE":
                if (data.down) addPoint(stats.value.down, new Date(data.down.timestamp), data.down.value / 1000000);
                if (data.up) addPoint(stats.value.up, new Date(data.up.timestamp), data.up.value / 1000000);
                if (data.latency) addPoint(stats.value.latency, new Date(data.latency.timestamp), data.latency.value);
                break;
            default:
                console.log("WS: unkown message", message.type);
        }
    };
}
connectWS();
</script>
