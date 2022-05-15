<template>
    <div
        class="fixed bottom-0 left-0 right-0 flex flex-col overflow-hidden md:visible md:relative md:py-10 px-4 md:px-0"
        :class="{ 'h-auto w-full md:w-1/3': show, 'h-8 w-full md:h-full md:w-12': !show }"
    >
        <div
            class="flex h-full w-full flex-col items-stretch justify-start rounded-t-3xl bg-starlink-card p-3 md:ml-4 md:flex-row md:rounded-l-3xl md:rounded-r-none md:p-6 md:pl-3"
        >
            <div class="flex flex-col items-center justify-center pb-3 md:h-full md:pb-0 md:pr-3">
                <button @click="toggle()" class="h-2 w-24 rounded-3xl bg-starlink-darkgray md:h-48 md:w-2"></button>
            </div>

            <div class="flex grow flex-col justify-between overflow-hidden">
                <SidebarDeviceList :devices="devices" />
                <SidebarOutages :outages="outages" />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
var show = ref(false);
defineProps<{
    devices: any;
    outages: Map<string, number>;
}>();

function toggle() {
    show.value = !show.value;
    console.log("Test");
}
</script>
