<script setup>
import TagsInput from 'vue3-tags-input'

const props = defineProps({
    name:String,
    modelValue:Array,
    placeholder:String,
})

const {value, errorMessage} = useField(()=> props.name, undefined,{
    syncVModel:true
})
const onTagChange = (newTags) =>{
    value.value = newTags
}
const style = computed(()=>{
    if (errorMessage.value){
        return 'error';
    }
    return '';
    }
)
</script>
<template>
    <div :class="style">
        <TagsInput @on-tags-changed="onTagChange" :add-tags-on-keys="[13,188]" :tags="value" :placeholder="props.placeholder" />
        <BaseErrorMessage v-if="errorMessage">{{errorMessage}}</BaseErrorMessage>
    </div>
</template>

<style lang="css">
.error .v3ti {
    border-color: red;
}
.error .v3ti--focus{
    border-color: red;
}

.v3ti{
    border-width: 2px;
    border-radius: 0.375rem;
    border-color: #e5e7eb;
    padding: 0.2rem;
}
.v3ti--focus{
    outline-width: 0px;
    border-color: black;
    box-shadow: none;
}
.v3ti .v3ti-tag{
    background-color: gray;
    font-size: small;
    padding: 0.2rem 0.5rem;
    border-radius: 0.375rem;
}
</style>