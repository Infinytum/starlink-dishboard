<template>
    <div class="grid grid-cols-1 gap-2 md:grid-cols-3">
        <NavbarStatusIndicator :status="status.ws" />
        <NavbarStatusIndicator :status="status.dishy" />
        <NavbarStatusIndicator :status="status.uplink" />
    </div>
</template>

<script setup lang="ts">
defineProps<{
    status: {
        ws: string,
        dishy: string,
        uplink: string,
    }
}>();
</script>