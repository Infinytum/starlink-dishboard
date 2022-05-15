<template>
    <div class="flex flex-col">
        <div class="relative">
            <div class="absolute inset-0 flex items-center" aria-hidden="true">
                <div class="w-full border-t border-gray-500"></div>
            </div>
            <div class="relative flex justify-center">
                <span class="bg-starlink-card px-4 text-gray-400">Outages</span>
            </div>
        </div>
        <SidebarOutagesChart :keys="[...outages.keys()]" :values="[...outages.values()]" />
    </div>
</template>

<script setup lang="ts">
defineProps<{
    outages: Map<string, number>;
}>();
</script>
