export function LoginApi(username,password){
    return useFetchAPI('/auth/login',{
        method:'post',
        body:{
            username,
            password,
        },
    });
}