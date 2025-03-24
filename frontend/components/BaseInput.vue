<script setup>

defineProps({
    name:String,
    placeholder:String,
    type:String,
})

const style = computed(()=>{
    return "w-full border-2 rounded-md px-3 py-2 focus:outline-none focus:border-blue-500 bg-white"
});

let value = ref("");
</script>
<template>
    <div class="mb-1 w-3/4 block">
        <input :class="style" v-model="value" :placeholder="placeholder" :type="type">
        <h2 class=" hidden">test</h2>
    </div>
</template>