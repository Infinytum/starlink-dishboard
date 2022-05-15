<template>
    <div class="grid grid-cols-2 gap-4 md:grid-cols-3 md:gap-10">
        <StatsGraph title="Latency" :data="latencyStats" />
        <StatsSingle title="Downloaded" :value="downloaded" unit="GB" />
        <StatsSingle title="Uploaded" :value="uploaded" unit="GB" />
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
    latencyStats: Map<string, number>;
    downloadStats: Map<string, number>;
    uploadStats: Map<string, number>;
}>();

var downloaded = computed(() => {
    if (!props.downloadStats) {
        return 0;
    }
    let values = [...props.downloadStats.values()];
    let total = 0;
    values.forEach(value => {
        total += value;
    });
    return total;
});
var uploaded = computed(() => {
    if (!props.uploadStats) {
        return 0;
    }
    let values = [...props.uploadStats.values()];
    let total = 0;
    values.forEach(value => {
        total += value;
    });
    return total;
});
</script>
