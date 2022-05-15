<template>
    <div class="grid grid-cols-2 gap-4 md:grid-cols-3 md:gap-10">
        <StatsGraph title="Latency" :keys="[...latencyStats.keys()]" :values="[...latencyStats.values()]" />
        <StatsSingle title="Downloaded" :value="downloaded" unit="GB" />
        <StatsSingle title="Uploaded" :value="uploaded" unit="GB" />
    </div>
</template>

<script setup lang="ts">
defineProps<{
    latencyStats: Map<string, number>;
    downloaded: number;
    uploaded: number;
}>();
</script>
