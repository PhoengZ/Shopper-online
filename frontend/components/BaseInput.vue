<script setup>

const props = defineProps({
    name:String,
    placeholder:String,
    modelvalue:String,
})

const {value,errorMessage} = useField(()=>props.name, undefined,{
    syncVModel:true,
    validateOnValueUpdate:false,
});

const style = computed(()=>{
    if (errorMessage.value){
        return "w-full border-2 rounded-md px-3 py-2 focus:outline-none border-red-500 bg-white";
    }
    return "w-full border-2 rounded-md px-3 py-2 focus:outline-none focus:border-blue-500 bg-white";
});

</script>
<template>
    <div class="mb-1 w-3/4 block">
        <input :class="style" v-model="value" :placeholder="placeholder" :type="name">
        <BaseErrorMessage v-if="errorMessage">{{errorMessage}}</BaseErrorMessage>
    </div>
</template>