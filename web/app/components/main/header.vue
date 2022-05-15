<template>
    <div class="flex items-center justify-center md:flex-row flex-col md:justify-between">
        <div class="grid grid-cols-2 gap-8">
            <div class="flex flex-col">
                <div class="flex items-center">
                    <div class="bg-starlink-lightgray mr-3 block h-2 w-10 rounded-3xl"></div>
                    <p class="font-sm text-starlink-lightgray">Download</p>
                </div>
                <p class="text-starlink-lightgray text-xs font-light">24 Mbit/s</p>
            </div>
            <div class="flex flex-col">
                <div class="flex items-center">
                    <div class="bg-starlink-darkgray mr-3 block h-2 w-10 rounded-3xl"></div>
                    <p class="font-sm text-starlink-darkgray">Upload</p>
                </div>
                <p class="text-starlink-darkgray text-xs font-light">3 Mbit/s</p>
            </div>
        </div>
        <div class="md-gap-3 grid grid-cols-3 gap-2 pt-8 md:pt-0">
            <button @click="select('today')" class="text-gray-500" :class="{'text-gray-300': selection == 'today'}">Today</button>
            <button @click="select('week')" class="text-gray-500" :class="{'text-gray-300': selection == 'week'}">Week</button>
            <button @click="select('month')" class="text-gray-500" :class="{'text-gray-300': selection == 'month'}">Month</button>
        </div>
    </div>
</template>

<script setup lang="ts">
var selection = ref("today");

function select(selected: string) {
    selection.value = selected;
}
</script>
