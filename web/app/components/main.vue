<template>
    <div class="grid grid-cols-1 gap-4">
        <MainHeader />
        <MainChart :keys="keys" :download="download" :upload="upload" />
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
