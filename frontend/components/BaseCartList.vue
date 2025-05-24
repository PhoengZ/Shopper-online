<script setup>
defineProps(
    {
        items:Array
    }
)
defineEmits(['add', 'remove'])

</script>
<template>
    <ul class=" flex flex-col gap-y-3 [&>li]:block p-2">
        <li v-for="item in items" :key="item.id" class=" w-full bg-white rounded-2xl drop-shadow-2xl">
            <BaseCartItem class=" font-bold grid grid-cols-5">
                <div class="flex items-center justify-center">
                    <BaseImage url="test"/>
                </div>
                <div class=" flex items-center justify-center text-center">
                    {{item.name}}
                </div>
                <div class=" flex items-center justify-center text-center">
                    {{item.quantity == null ? 1 :item.quantity}}
                </div>
                <div class=" flex items-center justify-center text-center">
                    <BaseButton size="small" theme="second" @click="$emit('add', item)">
                        <IconDropUp width="20" height="20"/>
                    </BaseButton>
                    <BaseButton size="small" theme="third" @click="$emit('remove', item)">
                        <IconDropDown width="20" height="20"/>
                    </BaseButton>
                </div>
                <div class=" flex items-center justify-center text-center">
                    {{item.price * (item.quantity == null ? 1 : item.quantity)}}
                </div>
            </BaseCartItem>
        </li>
    </ul>
</template>