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
export function getCartItem(id,token){
    return useFetchAPIMounted(`/cartList/${id}`,{
        method:'get',
        headers:{
            Authorization:`Bearer ${token}`
        }
    })
}
export function addItem(product,id,token){
    return useFetchAPIMounted('/addCartItem',{
        method:'post',
        headers:{
            Authorization:`Bearer ${token}`
        },
        body:{
            id:id,
            product:product
        }
    })
}
export function removeItem(userId, productId,token){
    return useFetchAPIMounted('/removeCartItem',{
        method:'patch',
        headers:{
            Authorization:`Bearer ${token}`
        },
        body:{
            userId:userId,
            itemId:productId
        }
    })
}

export function getProfile(id,token){
    return useFetchAPI(`/auth/getProfile/${id}`,{
        method:'get',
        headers:{
            Authorization:`Bearer ${token}`
        }
    })
}

export function updateProfile(id,profile,token){
    return useFetchAPIMounted('auth/updateProfile',{
        method:'patch',
        headers:{
            Authorization:`Bearer ${token}`
        },
        body:{
            id:id,
            profile:profile
        }
    })
}
