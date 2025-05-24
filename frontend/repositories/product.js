export function getProduct (){
    return useFetchAPI("/products",{
        method:'get'
    })
}
export function getProductByID (id){
    return useFetchAPI(`/product/${id}`,{
        method:'get'
    })
}