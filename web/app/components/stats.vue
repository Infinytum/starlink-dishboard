<template>
    <div class="grid grid-cols-2 gap-4 md:grid-cols-3 md:gap-10">
        <StatsGraph title="Latency" :data="latencyStats" />
        <StatsSingle title="Downloaded" :value="downloaded" unit="GB" />
        <StatsSingle title="Uploaded" :value="uploaded" unit="GB" />
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
    latencyStats: Map<Date, number>;
    downloadStats: Map<Date, number>;
    uploadStats: Map<Date, number>;
}>();

var downloaded = computed(() => {
    if (!props.downloadStats || props.downloadStats.size < 1) {
        return 0;
    }
    let values = [...props.downloadStats.values()];
    let total = 0;
    values.forEach(value => {
        total += value;
    });
    return total / 10000;
});
var uploaded = computed(() => {
    if (!props.uploadStats || props.uploadStats.size < 1) {
        return 0;
    }
    let values = [...props.uploadStats.values()];
    let total = 0;
    values.forEach(value => {
        total = total + value;
    });
    return total / 1000;
});
</script>
