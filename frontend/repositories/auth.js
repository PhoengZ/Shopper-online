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
    return useFetchAPIMounted('/auth/updateProfile',{
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

// export function updatePaying(id,value,token){
//     return useFetchAPIMounted('/auth/updatePaying',{
//         method:'patch',
//         headers:{
//             Authorization:`Bearer ${token}`
//         },
//         body:{
//             id:id,
//             coin:value,
//         }
//     })
// }

export function getStoreItem(id, token){
    return useFetchAPI(`/auth/getStoreItem/${id}`,{
        method:'get',
        headers:{
            Authorization:`Bearer ${token}`
        }
    })
}

export function addStoreItem(id, product, token){
    return useFetchAPIMounted('auth/addStoreItem',{
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

export function removeStoreItem(userID, productID, token){
    return useFetchAPIMounted('/auth/removeStoreItem',{
        method:'delete',
        headers:{
            Authorization:`Bearer ${token}`
        },
        body:{
            userID:userID,
            productID:productID
        }
    })
}

export function editStoreItem(userID, productID, product, token){
    return useFetchAPIMounted('/auth/editStoreItem',{
        method:'patch',
        headers:{
            Authorization:`Bearer ${token}`
        },
        body:{
            userID:userID,
            productID:productID,
            product:product
        }
    }
)
}