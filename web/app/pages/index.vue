<template>
    <div class="flex h-full flex-col md:flex-row">
        <div class="flex h-full grow flex-col justify-between p-4 pb-12 md:p-10">
            <Navbar />
            <Main :downloadStats="downloadStats" :uploadStats="uploadStats" />
            <Stats :latencyStats="latencyStats" :downloaded="downloaded" :uploaded="uploaded" />
        </div>

        <Sidebar :devices="devices" :outages="outages" />
    </div>
</template>

<script setup lang="ts">
var latencyStats = ref(
    new Map([
        ["0", 10],
        ["1", 41],
        ["2", 35],
        ["3", 51],
        ["4", 49],
        ["5", 62],
        ["6", 69],
    ])
);
var downloadStats = ref(
    new Map([
        ["0", 10],
        ["1", 41],
        ["2", 35],
        ["3", 51],
        ["4", 49],
        ["5", 62],
        ["6", 69],
    ])
);
var uploadStats = ref(
    new Map([
        ["0", 10],
        ["1", 41],
        ["2", 35],
        ["3", 51],
        ["4", 49],
        ["5", 62],
        ["6", 69],
    ])
);

var downloaded = ref(0);
var uploaded = ref(0);

var devices = ref([
    {
        type: "wired",
        name: "UDM-PRO",
        connection: "Connected - 1Gbit/s",
        vendor: "Ubiquiti Networks",
    },
    {
        type: "wireless",
        name: "Nintendo Switch",
        connection: "Wirelessly Connected - 5GHz",
        vendor: "Nintendo",
    },
    {
        type: "wireless",
        name: "Nintendo Switch",
        connection: "Wirelessly Connected - 5GHz",
        vendor: "Nintendo",
    },
    {
        type: "wireless",
        name: "Nintendo Switch",
        connection: "Wirelessly Connected - 5GHz",
        vendor: "Nintendo",
    },
    {
        type: "wireless",
        name: "Nintendo Switch",
        connection: "Wirelessly Connected - 5GHz",
        vendor: "Nintendo",
    },
    {
        type: "wireless",
        name: "Nintendo Switch",
        connection: "Wirelessly Connected - 5GHz",
        vendor: "Nintendo",
    },
    {
        type: "wireless",
        name: "Nintendo Switch",
        connection: "Wirelessly Connected - 5GHz",
        vendor: "Nintendo",
    },
    {
        type: "wireless",
        name: "Nintendo Switch",
        connection: "Wirelessly Connected - 5GHz",
        vendor: "Nintendo",
    },
    {
        type: "wireless",
        name: "Nintendo Switch",
        connection: "Wirelessly Connected - 5GHz",
        vendor: "Nintendo",
    },
    {
        type: "wireless",
        name: "Nintendo Switch",
        connection: "Wirelessly Connected - 5GHz",
        vendor: "Nintendo",
    },
]);
var outages = ref(
    new Map([
        ["0", 12],
        ["1", 3],
        ["2", 1],
        ["3", 6],
        ["4", 32],
        ["5", 3],
        ["6", 8],
    ])
);

// Initial Data Load
var initialDataLoad = new Map<string, number>();
for (let i = 0; i < 64; i++) {
    initialDataLoad.set(i.toString(), Math.floor(Math.random() * 100));
}
latencyStats.value = initialDataLoad;

// Do webscoket or rest call here
setInterval(() => {
    downloaded.value += Math.random() / 100;
    uploaded.value += Math.random() / 200;
}, 1000);
</script>
