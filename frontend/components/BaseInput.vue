<script setup>
const props = defineProps({
    name:String,
    placeholder:String,
    modelvalue:String,
    type:{
        type:String,
        default:'text'
    },
    width:{
        type:String,
        deafult:'w-3/4'
    },
})
const emit = defineEmits(['update:modelvalue']);
let value = ref(props.modelvalue); // fallback ถ้าไม่ใช้ vee-validate
let errorMessage = ref("");
if (props.name){
    const field = useField(() => props.name, undefined, {
        syncVModel: true,
        validateOnValueUpdate: false,
    });
    value = field.value;
    errorMessage = field.errorMessage;
}

watch(value,(newValue)=>{
    emit('update:modelvalue', newValue);
})
const style = computed(()=>{
    if (errorMessage.value){
        return "w-full border-2 rounded-md px-3 py-2 focus:outline-none border-red-500 bg-white";
    }
    return "w-full border-2 rounded-md px-3 py-2 focus:outline-none focus:border-blue-500 bg-white";
});

</script>
<template>
    <div class="mb-3  h-full" :class="props.width">
        <input :class="style" v-model="value" :placeholder="placeholder" :type="type" :accept="type === 'file' ? 'image/*' : ''"  class="h-full">
        <BaseErrorMessage v-if="errorMessage">{{errorMessage}}</BaseErrorMessage>
    </div>
</template>