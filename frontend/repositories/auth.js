export function LoginApi(username,password){
    return useFetchAPI('/auth/login',{
        method:'post',
        body:{
            username,
            password,
        },
    });
}

export function registerApi(username,password){
    return useFetchAPI('/auth/register',{
       method:'post',
       body:{
            username,
            password,
       }
    });
}
export function validateToken(token){
    return useFetchAPI('/validate',{
        method:'post',
        body:{
            token:token
        }
    })
}