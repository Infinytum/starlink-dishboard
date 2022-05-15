<template>
    <BarChart v-if="outages" v-bind="barChartProps" />
</template>

<script setup lang="ts">
import { BarChart, useBarChart } from "vue-chart-3";
import { ChartData, ChartOptions } from "chart.js";

const props = defineProps<{
    outages: Map<string, Map<string, number>>,
}>();

const chartData = computed<ChartData<"bar">>(() => {
    let labels = props.outages.size > 0 ? props.outages.values().next().value.keys() : [];
    let datasets = Array();
    props.outages.forEach((data: Map<string, number>, name: string) => {
        datasets.push({
            label: name,
            data: data.values(),
            backgroundColor: name == "obstruction" ? "#FFFFFF" : "#000000",
            borderColor: "rgba(0, 0, 0, 0.1)",
            borderWidth: 1,
        });
    });
    return {
        labels: labels,
        datasets: datasets,
    };
});

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
