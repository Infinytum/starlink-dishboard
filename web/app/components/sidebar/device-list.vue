<template>
    <ul class="pb-4 max-h-64 md:max-h-fit overflow-y-auto">
        <DeviceListItem v-for="device in devices" :device="device" />
    </ul>
</template>

<script setup lang="ts">
import DeviceListItem from './device-list-item.vue';
defineProps<{
    devices: any[];
}>();
</script>