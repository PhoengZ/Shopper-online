<script setup>
import {IsvalidUrl} from '../composables/Image.js';
import IconPlaceholder from './Icon/Placeholder.vue';
defineProps({
    url:String,
});
</script>
<template>
    <img v-if="IsvalidUrl(url)" :src="url">
    <slot v-else name="placeholder">
        <IconPlaceholder/>
    </slot>
</template>