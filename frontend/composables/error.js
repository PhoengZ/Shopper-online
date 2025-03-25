export const useCustomError = (error,callback)=>{
    if (error >= 400 || error < 500){
        return callback(error);
    }
    throw new Error(error);
};