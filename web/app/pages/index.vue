<template>
    <div class="flex h-full flex-col md:flex-row">
        <div class="flex h-full grow flex-col justify-between p-4 pb-12 md:p-10">
            <Navbar />
            <Main :downloadStats="data.downloadStats" :uploadStats="data.uploadStats" />
            <Stats
                :latencyStats="data.latencyStats"
                :downloadStats="data.downloadStats"
                :uploadStats="data.uploadStats"
            />
        </div>

        <Sidebar :devices="data.devices" :outages="data.outages" />
    </div>
</template>

<script setup lang="ts">
const LOOKBACK_LATENCY = 60;
const LOOKBACK_DOWNLOAD = 60;
const LOOKBACK_UPLOAD = 60;

const config = ref({
    selection: "live",
});
const data = ref({
    latencyStats: new Map<string, number>(),
    downloadStats: new Map<string, number>(),
    uploadStats: new Map<string, number>(),
    devices: [] as any[],
    outages: new Map<string, Map<string, number>>(),
});

function addLatencyPoint(value: number) {
    var last = data.value.latencyStats.size > 0 ? parseInt([...data.value.latencyStats.keys()].reverse()[0]) : 0;
    data.value.latencyStats.set((last + 1).toString(), value);
    if (data.value.latencyStats.size > LOOKBACK_UPLOAD) {
        data.value.latencyStats.delete(data.value.latencyStats.keys().next().value);
    }
}
function addDownloadPoint(value: number) {
    var last = data.value.downloadStats.size > 0 ? parseInt([...data.value.downloadStats.keys()].reverse()[0]) : 0;
    data.value.downloadStats.set((last + 1).toString(), value);
    if (data.value.downloadStats.size > LOOKBACK_UPLOAD) {
        data.value.downloadStats.delete(data.value.downloadStats.keys().next().value);
    }
}
function addUploadPoint(value: number) {
    var last = data.value.uploadStats.size > 0 ? parseInt([...data.value.uploadStats.keys()].reverse()[0]) : 0;
    data.value.uploadStats.set((last + 1).toString(), value);
    if (data.value.uploadStats.size > LOOKBACK_UPLOAD) {
        data.value.uploadStats.delete(data.value.uploadStats.keys().next().value);
    }
}

// Initial Data Load
var initialDataLoad = new Map<string, number>();
for (let i = 0; i < 60; i++) {
    addLatencyPoint(Math.floor(Math.random() * 100));
    addDownloadPoint(Math.floor(Math.random() * 100));
    addUploadPoint(Math.floor(Math.random() * 100));
}

// Do webscoket or rest call here
setInterval(() => {
    addLatencyPoint(Math.floor(Math.random() * 100));
    addDownloadPoint(Math.floor(Math.random() * 100));
    addUploadPoint(Math.floor(Math.random() * 100));
}, 1000);
</script>
