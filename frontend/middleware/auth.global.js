import { useAuthStore } from "~/Stores/auth";

export default defineNuxtRouteMiddleware((to,from)=>{
    const token = useCookie('token');
    const user = useAuthStore();
    if (token.value){
        user.setUser(token.value);
    }
})