<template>
    <button @click="$emit('click')" class="text-gray-500 p-2" :class="{ 'text-gray-300': active }">{{title}}</button>
</template>

<script setup lang="ts">
defineProps<{
    title: string;
    active: boolean;
}>();
</script>
