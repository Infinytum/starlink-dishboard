<template>
    <div class="flex justify-center md:justify-start">
        <MainLegendElement name="Download" :value="download" unit="Mbit/s" color="text-starlink-lightgray" bg="bg-starlink-lightgray" side="left" />
        <MainLegendElement name="Upload" :value="upload" unit="Mbit/s" color="text-starlink-darkgray" bg="bg-starlink-darkgray" side="right" class="ml-4 md:ml-10" />
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
    downloadStats?: Map<string, number>;
    uploadStats?: Map<string, number>;
}>();

var download = computed(() => {
    if (!props.downloadStats) {
        return 0;
    }
    var values = [...props.downloadStats.values()];
    return values[values.length - 1];
});
var upload = computed(() => {
    if (!props.uploadStats) {
        return 0;
    }
    var values = [...props.uploadStats.values()];
    return values[values.length - 1];
});
</script>
