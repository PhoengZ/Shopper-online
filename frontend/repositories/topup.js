export function requestTransaction(userID, amount, token){
    return useFetchAPIMounted('/topup/request',{
        method:'post',
        headers:{
            Authorization:`Bearer ${token}`
        },
        body:{
            userID:userID,
            amount:amount,
        }
    })
}