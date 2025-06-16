import { defineStore } from "pinia";
import { decodeToken } from "~/composables/Decodetoken";
import { GetUserDisplay, LoginApi, registerApi } from "~/repositories/auth";

export const useAuthStore = defineStore('auth',()=>{
    const user = ref({
        LoggedIn:false,
        token:'',
    });
    const isLoggedIn = ref(false);
    const userError = ref('');
    const userID = ref('');
    const Username = ref('');
    const coin = ref(0);
    const token = computed(()=>user.value.token);
    const canEdit = computed(()=>user.value.LoggedIn);
    async function Login(username,password) {
        isLoggedIn.value = true;
        const {data:response, error, status} = await LoginApi(username,password);
        if (status.value === 'error'){
            userError.value = error.value;
            return false;
        }
        user.value.LoggedIn = true;
        user.value.token = response.value.token;

        const token = useCookie('token');
        token.value = response.value.token;
        isLoggedIn.value = false;
        userError.value = '';
        Username.value = username;
        return true;
    }
    async function setUser(token) {
        user.value.LoggedIn = true;
        user.value.token = token;
        const item = decodeToken(token);
        Username.value = item.username;
        userID.value = item.userID;
    }
    async function getUserDisplay(){
        const {data: response, error, status} = await GetUserDisplay(userID.value,token.value)
        if (status === 'error'){
            userError.value = error.value;
            return false;
        }
        coin.value = response.value.coin
        return true
    }
    async function Logout() {
        user.value.LoggedIn = false;
        user.value.token = '';
        token.value = null;
        Username.value = '';
        userID.value = '';
    }
    async function Register(username,password){
        const {data:response, error, status}  = await registerApi(username,password);
        if (status === 'error'){
            useCustomError(error.value,(error)=>{
                userError.value = error.value;
            })
        }
        return true;
    }
    return {
        user,
        userError,
        userID,
        coin,
        token,
        canEdit,
        Username,
        setUser,
        Login,
        Logout,
        Register,
        getUserDisplay
    }
});