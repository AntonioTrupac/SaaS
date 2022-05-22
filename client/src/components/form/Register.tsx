import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';

import Button from '@/components/buttons/Button';
import { registerValidationSchema } from '@/components/form/validation';

import Form from './Form';

type IFormInput = {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
  country: string;
  city: string;
};

const Register = () => {
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
    // eslint-disable-next-line no-console
    console.log(data);
    reset();
  };

  return (
    <Form className='flex flex-col' onSubmit={handleSubmit(onSubmit)}>
      <input type='text' {...register('firstName')} placeholder='First Name' />
      <p>{errors.firstName?.message}</p>
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
