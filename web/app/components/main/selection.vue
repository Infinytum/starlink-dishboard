<template>
    <div class="flex justify-center md:justify-end items-center h-full">
        <MainSelectionButton title="Live" :active="selection == 'live'" @click="select('live')" />
        <MainSelectionButton title="Today" :active="selection == 'today'" @click="select('today')" />
        <MainSelectionButton title="Week" :active="selection == 'week'" @click="select('week')" />
        <MainSelectionButton title="Month" :active="selection == 'month'" @click="select('month')" />
    </div>
</template>

<script setup lang="ts">
var selection = ref("live");

function select(selected: string) {
    selection.value = selected;
}
</script>
