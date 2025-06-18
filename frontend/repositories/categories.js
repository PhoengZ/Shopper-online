export function getCategories(){
    return useFetchAPI("api/categories",{
        method:'get'
    })
}