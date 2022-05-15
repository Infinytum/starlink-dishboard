<template>
    <LineChart v-bind="lineChartProps" />
</template>

<script setup lang="ts">
import { LineChart, useLineChart } from "vue-chart-3";
import { ChartData, ChartOptions } from "chart.js";

const props = defineProps<{
    keys: string[];
    download: number[];
    upload: number[];
}>();

const chartData = computed<ChartData<"line">>(() => ({
    labels: props.keys,
    datasets: [
        {
            data: props.download,
            borderColor: "#EFEFEF",
            borderWidth: 5,
            pointRadius: 0,
            tension: 0.4,
        },
        {
            data: props.upload,
            borderColor: "#5D5D5D",
            borderWidth: 5,
            pointRadius: 0,
            tension: 0.4,
        },
    ],
}));
const options = computed<ChartOptions<"line">>(() => ({
    type: "line",
    responsive: true,
    interaction: {
        intersect: false,
    },
    aspectRatio: 2,
    animation: {
        duration: 0,
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
const { lineChartProps } = useLineChart({
    chartData,
    options,
});
</script>
