<script setup lang="ts">
import { onMounted } from 'vue';
import { state } from './lib/socket';


</script>

<template>
{{ state.connected }}
</template>
