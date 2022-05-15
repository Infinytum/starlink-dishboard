<template>
    <div class="col-span-2 flex flex-col rounded-2xl bg-starlink-card md:col-span-1">
        <div class="flex flex-row p-3 pb-0 md:p-4">
            <p class="font-bold text-white">{{ title }}</p>
            <p v-if="data" class="pl-2 font-light text-white">{{ last.toFixed(2) }}ms</p>
        </div>
        <LineChart v-if="data" v-bind="lineChartProps" />
    </div>
</template>

<script setup lang="ts">
import "chartjs-adapter-date-fns";
import { LineChart, useLineChart } from "vue-chart-3";
import { ChartData, ChartOptions } from "chart.js";

const props = defineProps<{
    title: string;
    data: Map<Date, number>;
}>();

const last = computed(() => {
    if (!props.data || props.data.size < 1) {
        return 0;
    }
    let values = Array.from(props.data.values());
    return values[values.length - 1];
});

const chartData = computed<ChartData<"line">>(() => ({
    labels: Array.from(props.data.keys()),
    datasets: [
        {
            data: Array.from(props.data.values()),
            borderColor: "#5D5D5D",
            borderWidth: 3,
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
    aspectRatio: 4,
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
