<template>
    <div class="flex flex-col justify-between rounded-2xl bg-starlink-card">
        <div class="flex flex-row p-3 pb-0 md:p-4">
            <p class="font-bold text-white">{{ title }}</p>
        </div>
        <div class="flex grow flex-col justify-end p-3 pt-0 md:p-4">
            <p class="font-extralight text-white">
                <strong class="text-3xl font-medium">{{ value.toFixed(2) }}</strong> {{ unit }}
            </p>
        </div>
    </div>
</template>

<script setup lang="ts">
defineProps<{
    title: string;
    value: number;
    unit: string;
}>();
</script>
