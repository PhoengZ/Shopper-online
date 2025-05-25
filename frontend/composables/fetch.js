export const useFetchAPI = (request,object)=>{
    const config = useRuntimeConfig();
    return useFetch(request,{baseURL:config.public.baseApiUrl,...object})
}
export const useFetchAPIMounted = async (request,object) =>{
    const config = useRuntimeConfig();
    return $fetch(request, {baseURL: config.public.baseApiUrl,...object});
}