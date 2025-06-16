export const useFetchAPI = (request,object)=>{
    const config = useRuntimeConfig();
    return useFetch(request,{baseURL:config.public.baseApiUrl,...object})
}
export const useFetchAPIMounted = async (request,object) =>{
    const config = useRuntimeConfig();
    try{
        const response = $fetch(request, {baseURL: config.public.baseApiUrl,...object});
        return response;
    }catch(err){
        if (err.response && err.response._data){
            const backendErrorData = err.response._data;
            const customError = new Error(
                backendErrorData.error || backendErrorData.message || `An error occurred with status ${err.response.status}`
            )
            customError.statusCode = err.response.status;
            customError.data = backendErrorData;
            throw customError;
        }
        throw err
    }
}