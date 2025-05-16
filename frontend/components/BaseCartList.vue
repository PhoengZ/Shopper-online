<script setup>

defineProps(
    {
        items:Array
    }
)
</script>
<template>
    <ul class=" flex flex-col gap-y-3 [&>li]:block p-2">
        <li v-for="item in items" :key="item.id" class=" w-full bg-white rounded-2xl drop-shadow-2xl">
            <BaseCartItem class=" font-bold grid grid-cols-4">
                <div class="flex items-center justify-center">
                    <BaseImage url="test"/>
                </div>
                <div class=" flex items-center justify-center text-center">
                    {{item.name}}
                </div>
                <div class=" flex items-center justify-center text-center">
                    {{item.quantity}}
                </div>
                <div class=" flex items-center justify-center text-center">
                    {{item.price * item.quantity}}
                </div>
            </BaseCartItem>
        </li>
    </ul>
</template>