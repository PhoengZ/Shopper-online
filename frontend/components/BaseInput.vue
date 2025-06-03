<script setup>
const props = defineProps({
    name:String,
    placeholder:String,
    modelvalue:{
        type:[String,Number,Array,File]
    },
    type:{
        type:String,
        default:'text'
    },
    width:{
        type:String,
        deafult:'w-3/4'
    },
    height:{
        type:String,
        default:'h-full'
    },
    Update:{
        type:Boolean,
        default:false
    }
})
const emit = defineEmits(['update:modelvalue']);
let value = ref(props.modelvalue); // fallback ถ้าไม่ใช้ vee-validate
let errorMessage = ref("");
if (props.name){
    const field = useField(() => props.name, undefined, {
        syncVModel: true,
        validateOnValueUpdate: props.Update,
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
const handleChange = (e) =>{
    const file = e.target.files[0]
    value.value = file
    emit('update:modelvalue', file)
}
</script>
<template>
    <div class="mb-3" :class="props.width,props.height">
        <input v-if="type === 'file'" type="file" :class="style" :placeholder="placeholder" @change="handleChange" class="h-full" :accept="'image/*'">
        <input v-else :class="style"  v-model="value" :placeholder="placeholder" :type="type" class="h-full"/>
        <BaseErrorMessage v-if="errorMessage">{{errorMessage}}</BaseErrorMessage>
    </div>
</template>