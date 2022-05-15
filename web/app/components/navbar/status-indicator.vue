<template>
    <div class="h-2 w-2 rounded-md md:h-3 md:w-3" v-bind:class = "status == 'online' ? 'bg-green-500' : status == 'offline' ? 'bg-red-500' : 'bg-starlink-lightgray'"></div>
</template>

<script setup lang="ts">
defineProps<{
    status?: string;
}>();
</script>
