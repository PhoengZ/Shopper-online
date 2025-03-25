export const useFetchAPI = (request,object)=>{
    const config = useRuntimeConfig();
    // console.log("Base API URL:", config.public.baseApiUrl);
    return useFetch(request,{baseURL:config.public.baseApiUrl,...object})
}