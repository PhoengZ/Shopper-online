<script setup>
import {IsvalidUrl} from '../composables/Image.js';
import IconPlaceholder from './Icon/Placeholder.vue';
defineOptions({
    inheritAttrs:false,
})

defineProps({
    url:String,
});
</script>
<template>
    <img v-if="IsvalidUrl(url)" :src="url" v-bind="$attrs" class=" rounded-2xl">
    <slot v-else name="placeholder">
        <IconPlaceholder/>
    </slot>
</template>