export function getProduct (page, limit){
    return useFetchAPIMounted(`api/products?page=${page}&limit=${limit}`,{
        method:'get'
    })
}
export function getProductByID (id){
    return useFetchAPI(`api/products/${id}`,{
        method:'get'
    })
}
export function getProductBySearching (data,page,limit){
    return useFetchAPIMounted(`api/products?name=${data.name}&price=${data.price}&category=${data.category.join(",")}&page=${page}&limit=${limit}`,{
        method:'get',
    })
}
export function getStoreItem(id, token){
    return useFetchAPI(`api/auth/getStoreItem/${id}`,{
        method:'get',
        headers:{
            Authorization:`Bearer ${token}`
        }
    })
}

export function addStoreItem(product, token){
    return useFetchAPIMounted('api/auth/addStoreItem',{
        method:'post',
        headers:{
            Authorization:`Bearer ${token}`
        },
        body:product
    })
}
export function removeStoreItem(productID, token){
    return useFetchAPIMounted('api/auth/removeStoreItem',{
        method:'delete',
        headers:{
            Authorization:`Bearer ${token}`
        },
        body:{
            productID:productID
        }
    })
}

export function editStoreItem(product, token){
    return useFetchAPIMounted('api/auth/editStoreItem',{
        method:'patch',
        headers:{
            Authorization:`Bearer ${token}`
        },
        body:product
    })
}

export function getCountProduct(){
    return useFetchAPI('api/product/count',{
        method:'get'
    })
}