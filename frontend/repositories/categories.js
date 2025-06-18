export function getCategories(){
    return useFetchAPI("/categories",{
        method:'get'
    })
}