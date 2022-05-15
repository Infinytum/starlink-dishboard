<template>
    <div class="grid grid-cols-1 gap-4 md:grid-cols-2 h-full md:h-auto">
        <div class="col-1 order-3 md:order-1">
            <MainLegend :download-stats="downloadStats" :upload-stats="uploadStats" />
        </div>
        <div class="col-2 order-1 md:order-2">
            <MainSelection />
        </div>
        <div class="md:col-span-2 order-2 md:order-3">
            <MainChart :keys="keys" :download="download" :upload="upload" />
        </div>
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
    downloadStats: Map<string, number>;
    uploadStats: Map<string, number>;
}>();

const keys = computed(() => {
    let k = [...props.downloadStats.keys()];
    let k2 = [];
    for (let i = 0; i < k.length; i++) {
        k2.push(k[i].toString());
    }
    return k2;
});
const download = computed(() => {
    return [...props.downloadStats.values()];
});
const upload = computed(() => {
    return [...props.uploadStats.values()];
});
</script>
