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