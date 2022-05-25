import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';

import useRegisterMutation from '@/hooks/useRegisterMutation';

import { RegisterVariables } from '@/graphql';
import Button from '@/ui/buttons/Button';
import { registerValidationSchema } from '@/ui/form/validation';
import TextField from '@/ui/textField/TextField';

import Form from './Form';

export type IFormInput = {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
  country: string;
  city: string;
};

const Register = () => {
  const mutation = useRegisterMutation();
  const { mutate } = mutation;
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<IFormInput>({
    defaultValues: {
      firstName: '',
      lastName: '',
      email: '',
      password: '',
      city: '',
      country: '',
    },
    mode: 'onBlur',
    resolver: yupResolver(registerValidationSchema),
  });

  const onSubmit: SubmitHandler<IFormInput> = (data) => {
    const mappedData: RegisterVariables = {
      input: {
        firstName: data.firstName,
        lastName: data.lastName,
        email: data.email,
        age: 18,
        password: data.password,
        phone: '094-4459-3434',
        address: [
          {
            city: data.city,
            addressLine: 'Donja Svarca',
            country: data.country,
            postalCode: 47250,
          },
        ],
      },
    };

    mutate(mappedData);
    reset();
  };

  return (
    <Form className='flex flex-col' onSubmit={handleSubmit(onSubmit)}>
      <TextField
        type='text'
        register={register}
        inputFieldName='firstName'
        placeholder='First Name'
        disabled={false}
      />
      {/*<input type='text' {...register('firstName')} placeholder='First Name' />*/}
      {/*<p>{errors.firstName?.message}</p>*/}
      <input type='text' {...register('lastName')} placeholder='Last Name' />
      <p>{errors.lastName?.message}</p>
      <input type='email' {...register('email')} placeholder='Email' />
      <p>{errors.email?.message}</p>
      <input type='password' {...register('password')} placeholder='Password' />
      <p>{errors.password?.message}</p>
      {/*<input type='text' /> TODO: GET ALL THE COUNTRIES API AND PREPARE A SELECT FIELD WITH AN AUTOCOMPLETE */}
      <input type='text' {...register('country')} placeholder='Country' />
      <p>{errors.country?.message}</p>
      <input type='text' {...register('city')} placeholder='City' />
      <p>{errors.city?.message}</p>

      <Button variant='primary' type='submit'>
        Submit
      </Button>
    </Form>
  );
};

export default Register;
