export function getProduct (){
    return useFetchAPI("/products",{
        method:'get'
    })
}
export function getProductByID (id){
    return useFetchAPI(`/products/${id}`,{
        method:'get'
    })
}
export function getProductBySearching (data){
    return useFetchAPIMounted(`/products?name=${data.name}&price=${data.price}&category=${data.category.join(",")}`,{
        method:'get',
    })
}
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