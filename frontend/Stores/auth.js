import { defineStore } from "pinia";
import { LoginApi, registerApi } from "~/repositories/auth";

export const useAuthStore = defineStore('auth',()=>{
    const user = ref({
        LoggedIn:false,
        token:'',
    });
    const isLoggedIn = ref(false);
    const userError = ref('');
    const Username = ref('');
    const token = computed(()=>user.value.token);
    const canEdit = computed(()=>user.value.LoggedIn);
    async function Login(username,password) {
        isLoggedIn.value = true;
        const {data:response, error, status} = await LoginApi(username,password);
        if (status.value === 'error'){
            useCustomError(error.value,(error)=>{
                userError.value = error.value;
            })
        }
        user.value.LoggedIn = true;
        user.value.token = response.value.token;

        const token = useCookie('token');
        token.value = username;
        isLoggedIn.value = false;
        userError.value = '';
        Username.value = username;
        return true;
    }
    async function setUser(token) {
        user.value.LoggedIn = true;
        user.value.token = token;
    }
    async function Logout() {
        user.value.LoggedIn = false;
        user.value.token = '';
        token.value = null;
        Username.value = '';
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
        token,
        canEdit,
        Username,
        setUser,
        Login,
        Logout,
        Register
    }
});