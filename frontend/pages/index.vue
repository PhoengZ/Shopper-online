<script setup>
import { useAuthStore } from '~/Stores/auth';
import { addItem, getCartItem, removeItem, updateProfile, validateToken } from '~/repositories/auth';
import { getProduct, getProductBySearching } from '~/repositories/product';
import { getCategories } from '../repositories/categories';
import BasePagelist from '~/components/ฺBasePagelist.vue';
const {$toast} = useNuxtApp()
useHead({
    title:"Shopper-Online"
})
let showList = ref(false);
let Item = ref([]);
const token = useCookie('token');
const user = useAuthStore();
const name = ref('');
const userID = ref('');
let showSetting = ref(false);
const totalPage = ref(0)
const nowPage = ref(1)
const limit = ref(5)
const pd = ref([]);
const filter = ref({
    name: '',
    price: '',
    category: [],
});
try{
    const {object,total} = await getProduct(nowPage.value,limit.value);
    pd.value = object
    totalPage.value = Math.ceil(Number(total)/limit.value)
}catch(err){
    console.error(err.data.error)
}
const choiceItem = ref([]);
const checkUser = await user.getUserDisplay()
const coin = ref(0)
if (checkUser){
    coin.value = user.coin
}
const {data:response, error:er, status} = await getCategories();
if (status.value === 'error'){
    console.error('Failed to fetch categories', er.value);
    choiceItem.value = [];
    $toast.error('ไม่สามารถเรียกดูประเภทสินค้าได้',{
        description:'รายละเอียด: ' + er.value,
    })
} else {
    choiceItem.value = response.value.categories;
    if (choiceItem.value[0] != "Price"){
        choiceItem.value.unshift("Price")
    }
} 
const {error:err, data: validateData} = await validateToken(token.value)
const isValidToken = computed(() => validateData.value?.message === 'Valid')
if (isValidToken){
    name.value = user.Username
    userID.value = user.userID;
}
onMounted(()=>{
    if (isValidToken){
        name.value = user.Username;
        userID.value = user.userID;
    }
})
const checkLogout = ()=>{
    if (isValidToken.value){
        token.value = null;
        name.value = '';
        user.Username = '';
        userID.value = '';
        navigateTo('/login');
        return 
    }
}

const checkAuth = ()=>{
    if (!isValidToken.value){
        navigateTo('/login');
        return 
    }
    showSetting.value = !showSetting.value
};

const checkItem = async ()=>{
    if (!isValidToken.value){
        navigateTo('/login');
        return 
    }
    if (showList.value){
        showList.value = !showList.value
        return 
    }
    try{
        const {products} = await getCartItem(userID.value,token.value);
        Item.value = products
        showList.value = !showList.value;
        
    }catch(err){
        if (err.response?.status == 401){
            navigateTo('/login')
        }else{
            console.error("Unexpected error:", err);
            $toast.error('ไม่สามารถเรียกดูตะกร้าสินค้าได้',{
                description:'รายละเอียด: '+err.data.error
            })
        }
    }
}
const Adding = async (item)=>{
    if (!isValidToken.value){
        navigateTo('/login');
        return 
    }
    try{
        const {message} = await addItem(item,userID.value,token.value)
        const {products} = await getCartItem(userID.value,token.value)
        Item.value = products
        $toast.success(message, {
            style: {
                background: 'green',
                color: 'white',
            },
        });
    }catch(err){
        if (err.response?.status == 401){
            navigateTo('/login')
        }else{
            console.error(err.data.error)
            $toast.error('ไม่สามารถเพิ่มสินค้าลงตะกร้าได้', {
                description: 'รายละเอียด: ' + err.data.error, 
            });
        }
    }
} 
const Cancle = async (item)=>{
    if (!isValidToken.value){
        navigateTo('/login');
        return
    }
    try{
        const {message} = await removeItem(userID.value,item.id,token.value)
        const {products} = await getCartItem(userID.value,token.value)
        Item.value = products
        $toast.success(message, {
            style: {
                background: 'green',
                color: 'white',
            },
        });
    }catch(err){
        if (err.response?.status == 401){
            navigateTo('/login')
        }else{
            console.error("Unexpected error:", err.data.error)
            $toast.error('ไม่สามารถยกเลิกรายการสินค้าได้',{
                description:'รายละเอียด: '+err.data.error
            })
        }
    }
}
const SearchItem = async (block)=>{
    try{
        const {object,total} = await getProductBySearching(block,nowPage.value,limit.value)
        filter.value = block
        pd.value = object
        totalPage.value = Math.ceil(Number(total)/limit.value)
        nowPage.value = 1
    }catch(err){
        console.error(err)
        $toast.error('ไม่สามารถค้นหารายการสินค้าได้',{
            description:'รายละเอียด: '+err.data.error
        })
    }
}
const handleProfile = ()=>{
    if (!isValidToken.value){
        navigateTo('/login');
        return 
    }
    navigateTo('/profile');
}
const handleBuyItem = async(item,totalPrice) =>{
    if (!isValidToken.value){
        navigateTo('/login')
        return
    }
    const object = {
        "password":"",
        "address":"",
        "coin":-totalPrice,
        "history":item
    }
    try{
        if (totalPrice == 0){
            $toast.error('ไม่สามารถสั่งซื้อสินค้าได้',{
                description:'รายละเอียด: ไม่มีการเลือกสินค้า'
            })
            return
        }
        const {message} = await updateProfile(userID.value,object,token.value)
        const {products} = await getCartItem(userID.value, token.value)
        Item.value = products
        $toast.success(message, {
            style: {
                background: 'green',
                color: 'white',
            },
        });
    }catch(err){
        console.error(err.data.error);
        $toast.error('ไม่สามารถสั่งซื้อสินค้าได้',{
            description:'รายละเอียด: '+err.data.error
        })
    }
}
const handleOutside = ()=>{
    showList.value = false;
}
const handleProduct = ()=>{
    if (!isValidToken){
        navigateTo('/login')
    }
    navigateTo('/selling')
}   
const handleTopup = ()=>{
    if (!isValidToken.value){
        navigateTo('/login')
        return
    }
    navigateTo('/topup');
}
const changePage = async(page)=>{
    try{
        const {object} = await getProductBySearching(filter.value,page,limit.value)
        pd.value = object
        nowPage.value = page
    }catch(err){
        console.error(err.data.error)
        $toast.error('ไม่สามารถเปลี่ยนหน้าต่างสินค้าได้',{
            description:'รายละเอียด: '+err.data.error
        })
    }
}
</script>

<template>
    <AppHeader 
        :coin="coin" 
        :choiceItem="choiceItem" 
        :username="name" 
        :isDrop="showSetting" 
        :openBlure="showList" 
        @topup="handleTopup" 
        @getproduct="handleProduct" 
        @profile="handleProfile" 
        @searchItem="SearchItem" 
        @logout="checkLogout" 
        @auth="checkAuth" 
        @checkItem="checkItem"
    />
    
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <section class="bg-white rounded-2xl shadow-sm border border-neutral-100 overflow-hidden min-h-[600px]" :class="showList ? 'blur-sm transition duration-300':''">
             <BaseCardList class="p-6" :product="pd" @buy="Adding" mode="main" />
        </section>
        
        <div class="mt-8 flex justify-center">
            <BasePagelist :totalPage="totalPage" :nowPage="nowPage" @changePage="changePage"/>
        </div>
    </div>

    <CartForm v-if="showList" :item="Item" v-click-outside="handleOutside" @buy="handleBuyItem" @add="Adding" @remove="Cancle"/>
</template>
