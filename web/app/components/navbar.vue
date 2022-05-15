<template>
    <div class="flex w-full items-center">
        <NavbarLogo />
        <div class="grow"></div>
        <NavbarStatus :status="status" />
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
