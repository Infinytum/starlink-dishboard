<template>
    <div class="flex flex-row">
        <div class="flex items-center" :class="class1">
            <div class="h-2 w-8 rounded-3xl" :class="bg"></div>
        </div>
        <div :class="class2">
            <p class="font-sm" :class="p2Classes">{{ name }}</p>
            <p class="text-xs font-light" :class="p2Classes" >
                {{ value.toFixed(1) }} {{ unit }}
            </p>
        </div>
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
    name: string;
    value: number;
    unit: string;
    color: string;
    bg: string;
    side: string;
}>();
const class1 = computed(() => {
    return props.side == "left" ? "order-2 md:order-1 ml-3 md:ml-0 mr-3" : "mr-3 md:mr-3";
})
const class2 = computed(() => {
    return props.side == "left" ? "order-1 md:order-2" : "";
})
const p2Classes = computed(() => {
    return props.color + " " + (props.side == 'left' ? 'text-right md:text-left' : '');
});
</script>
