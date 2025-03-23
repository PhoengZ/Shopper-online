<script setup>

let pd = ref([{name:'IPhone13',des:'The iPhone is a sleek, high-performance smartphone with advanced cameras, a powerful chip, and an intuitive iOS experience.'},
{name:'IPhone14',des:'The iPhone is a sleek, high-performance smartphone with advanced cameras, a powerful chip, and an intuitive iOS experience.'},
{name:'IPhone15',des:'The iPhone is a sleek, high-performance smartphone with advanced cameras, a powerful chip, and an intuitive iOS experience.'}]);
useHead({
    title:"Shopper-online",
});


</script>

<template>
    <!-- นี้คือหน้าหลัก
    <BaseBadgeList :badges="['a', 'b']" />
    <BaseButton size="small" theme="third">
        <IconCart></IconCart>
        Hello
    </BaseButton>
    <BaseImage url=""/> -->
    <section>
        <!-- Part of showing product  -->
         <BaseCardList class="p-6" :product="pd"/>
    </section>
</template>