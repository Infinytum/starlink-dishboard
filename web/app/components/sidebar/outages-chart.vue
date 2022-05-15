<template>
    <BarChart v-bind="barChartProps" />
</template>

<script setup lang="ts">
import { BarChart, useBarChart } from "vue-chart-3";
import { ChartData, ChartOptions } from "chart.js";

const props = defineProps<{
    keys: string[];
    values: number[];
}>();

const chartData = computed<ChartData<"bar">>(() => ({
    labels: props.keys,
    datasets: [
        {
            data: props.values,
            backgroundColor: "#FFFFFF",
            borderColor: "transparent",
            borderRadius: 10,
        },
        {
            data: props.values,
            backgroundColor: "#5D5D5D",
            borderColor: "transparent",
            borderRadius: 10,
        },
    ],
}));
const options = computed<ChartOptions<"bar">>(() => ({
    type: "bar",
    responsive: true,
    interaction: {
        intersect: false,
    },
    aspectRatio: 1.75,
    animation: {
        duration: 0,
    },
    elements: {
        bar: {
            borderWidth: 4,
        },
    },
    scales: {
        x: {
            display: false,
        },
        y: {
            display: false,
        },
    },
    plugins: {
        title: {
            display: true,
        },
        legend: {
            display: false,
        },
    },
}));
const { barChartProps } = useBarChart({
    chartData,
    options,
});
</script>
