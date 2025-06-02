<script setup>
import { useCategoryValidationShema } from '~/composables/Validation';
import BaseInputTags from './BaseInputTags.vue';

const props = defineProps({
    categories:Array,
})
const emit = defineEmits(['submit','cancel'])
const {handleSubmit, isSubmitting} = useForm({
    validationSchema:useCategoryValidationShema,
    validateOnInput:true,
    keepValuesOnUnmount:true,
    initialValues:{
        categories: props.categories || []
    }
})
const onSubmit = handleSubmit(values  =>{
    emit('submit', values.categories)
})
</script>
<template>
    <form @submit.prevent="onSubmit" class=" flex gap-4 flex-col">
        <BaseInputTags name="categories" placeholder="Category (comma or enter separated)"/>
        <BaseButton type="submit" size="small" theme="first" :disabled="isSubmitting">{{isSubmitting ? "Saving Category":"Save Category"}}</BaseButton>
    </form>
</template>