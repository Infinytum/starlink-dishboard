<template>
    <li class="flex justify-between border-b-2 border-b-gray-500 last:border-b-0 py-2">
        <div class="flex items-center">
            <GlobeIcon class="h-5 w-5" v-if="device.type == 'wired'" />
            <WifiIcon class="h-5 w-5" v-if="device.type == 'wireless'" />
            <span class="px-4 py-2 font-bold text-white">{{ device.name }}</span>
        </div>

        <div class="flex flex-col content-center items-end text-right">
            <span class="text-sm text-green-500">{{ device.connection }}</span>
            <span class="text-xs text-gray-400">{{ device.vendor }}</span>
        </div>
    </li>
</template>

<script setup lang="ts">
import { GlobeIcon, WifiIcon } from "@heroicons/vue/outline";

defineProps<{
    device: any;
}>();
</script>
