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
export function getProductBySearching (data){
    return useFetchAPI("/product",{
        method:'get',
        body:{
            data:data
        }
    })
}