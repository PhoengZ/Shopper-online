import * as yup from  'yup';
export const useLoginValidationSchema = ()=>yup.object({
    username: yup.string().min(3).required().label('Username'),
    password: yup.string().required().label('Password'),
});

export const useRegisterValidationSchema = ()=>yup.object({
    username:yup.string().min(6).required().label('Username'),
    password_1:yup.string().min(6).required().label('Password'),
    password_2:yup.string().oneOf([yup.ref('password_1')],"Password must match").required().label('Confirm Passowrd')
});

export const useCategoryValidationShema = ()=>yup.object({
    categories: yup.array().min(1).required().of(yup.string().required().min(1)).label('categories'), 
})
export const useProductValidationSchema = ()=>yup.object({
    name: yup.string().min(3).required().label("Name"),
    description: yup.string().min(3).required().label("Description"),
    price: yup.number().min(1).required().label("Price"),
    quantity: yup.number().min(1).required().label("Quantity"),
})