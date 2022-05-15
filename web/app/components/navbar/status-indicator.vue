<template>
    <div
        class="h-2 w-2 rounded-sm md:h-3 md:w-3"
        v-bind:class="
            status == 'online' ? 'bg-green-400' : status == 'offline' ? 'bg-red-400' : 'bg-starlink-lightgray'
        "
    ></div>
</template>

<script setup lang="ts">
defineProps<{
    status: string;
}>();
</script>
