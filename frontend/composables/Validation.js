import * as yup from  'yup';
export const useLoginValidationSchema = ()=>yup.object({
    username: yup.string().min(3).required().label('Username'),
    password: yup.string().required().label('Password'),
});