<template>
    <LineChart :chartData="chartData" :options="options" />
</template>

<script setup lang="ts">
import "chartjs-adapter-date-fns";
import { LineChart, useLineChart } from "vue-chart-3";
import { ChartData, ChartOptions } from "chart.js";

const props = defineProps<{
    keys: Date[];
    download: number[];
    upload: number[];
}>();

const chartData = computed<ChartData<"line">>(() => ({
    labels: props.keys,
    datasets: [
        {
            label: "Download",
            data: props.download,
            borderColor: "#EFEFEF",
            borderWidth: 4,
            pointRadius: 0,
            tension: 0.4,
        },
        {
            label: "Upload",
            data: props.upload,
            borderColor: "#5D5D5D",
            borderWidth: 4,
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
    aspectRatio: 1,
    animation: {
        duration: 0,
    },
    scales: {
        x: {
            display: false,
            type: 'time',
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
