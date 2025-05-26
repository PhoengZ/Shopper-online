<script setup>
definePageMeta({
    layout:false
})
useHead({
    title:"Profile"
})
const isEditAddress = ref(false);
</script>
<template>
    <div class=" bg-gray-200 h-full w-screen mx-0 my-0 flex  py-20">
        <section class=" bg-white rounded-2xl h-full w-2/4 mx-auto my-auto drop-shadow-2xl flex flex-col py-7.5 gap-10 px-7.5">
            <h1 class=" mx-auto font-bold text-5xl">PROFILE</h1>
            <ul class=" grid grid-cols-2 gap-4 [&> li]: flex flex-col gap-10">
                <li>
                    <BaseImage url="https://f.ptcdn.info/090/014/000/1388837662-pantiptalk-o.png" 
                    class=" w-32 h-32 rounded-full border-2 mx-auto my-auto duration-300 transform  hover:scale-105"/>
                </li>
                <li>
                    <h2 class=" mx-auto my-auto text-center font-bold text-2xl mb-5">Delivery address</h2>
                    <textarea class="w-full h-20 border-2 border-gray-300 rounded-2xl px-3 py-3 resize-none" placeholder="Enter your delivery address here..." :disabled="!isEditAddress"></textarea>
                </li>
                <li>
                    <ul class=" grid grid-cols-2 gap-5 [&> li]: flex flex-row">
                        <li>
                            <h3 class=" text-right font-bold text-1xl my-auto mb-3">Username:</h3>
                        </li>
                        <li>
                            <input type="text" class="w-full h-10 border-2 border-gray-300 rounded-2xl px-3 py-2" 
                            placeholder="Your username" 
                            disabled
                            value="ChangeYourUsernameHere" />
                        </li>
                        <li>
                            <h3 class=" text-right font-bold text-1xlmy-auto">Password:</h3>
                        </li>
                        <li>
                            <input type="password" class="w-full h-10 border-2 border-gray-300 rounded-2xl px-3 py-2" 
                            placeholder="Your password" 
                            disabled
                            value="ChangePasswordhere" />
                        </li>
                    </ul>
                </li>
                <li>
                    <div class=" h-full flex flex-row my-auto mx-auto">
                        <BaseButton size="small" theme="second" class=" w-30 mx-auto my-auto" @click="isEditAddress = !isEditAddress">
                            {{ isEditAddress ? 'Save Address' : 'Edit Address' }}
                        </BaseButton>
                        <BaseButton size="small" theme="third" class=" w-30 mx-auto my-auto" >Forget Password</BaseButton>
                    </div>
                    
                </li>
            </ul>
            <BaseButton size="large" theme="third" class=" w-40 mx-auto my-auto">Sign-Out</BaseButton>
        </section>
    </div>
    
</template>