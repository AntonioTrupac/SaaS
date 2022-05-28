import { yupResolver } from '@hookform/resolvers/yup';
import { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';

import useRegisterMutation from '@/hooks/mutations/useRegisterMutation';

import { registerValidationSchema } from '@/components/form/validation';
import Button from '@/components/ui/buttons/Button';
import Label from '@/components/ui/label/Label';
import TextField from '@/components/ui/textField/TextField';

import { RegisterVariables } from '@/graphql';

import Form from './Form';

export type IFormInput = {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
  country: string;
  city: string;
};

type State = {
  isLoading: boolean;
  isError: boolean;
  isSuccess: boolean;
};

const Register = () => {
  const mutation = useRegisterMutation();
  const { mutate, isLoading, isError, isSuccess } = mutation;
  const [mutationStates] = useState<State>({
    isLoading,
    isError,
    isSuccess,
  });

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
      <div className='my-1.5 w-full'>
        <Label htmlFor='firstName' name='First Name' />
        <TextField
          type='text'
          register={register}
          inputFieldName='firstName'
          placeholder='First Name'
          states={mutationStates}
          errors={errors}
        />
      </div>

      <div className='my-1.5 w-full'>
        <Label htmlFor='lastName' name='Last Name' />
        <TextField
          type='text'
          register={register}
          inputFieldName='lastName'
          placeholder='Last Name'
          states={mutationStates}
          errors={errors}
        />
      </div>

      <div className='my-1.5 w-full'>
        <Label htmlFor='email' name='Email' />
        <TextField
          type='email'
          register={register}
          inputFieldName='email'
          placeholder='Email'
          states={mutationStates}
          errors={errors}
        />
      </div>

      <div className='my-1.5 w-full'>
        <Label htmlFor='password' name='Password' />
        <TextField
          type='password'
          register={register}
          inputFieldName='password'
          placeholder='Password'
          states={mutationStates}
          errors={errors}
        />
      </div>

      <div className='my-1.5 w-full'>
        <Label htmlFor='country' name='Country' />
        <TextField
          type='text'
          register={register}
          inputFieldName='country'
          placeholder='Country'
          states={mutationStates}
          errors={errors}
        />
      </div>

      <div className='mt-1.5 w-full'>
        <Label htmlFor='city' name='City' />
        <TextField
          type='text'
          register={register}
          inputFieldName='city'
          placeholder='City'
          states={mutationStates}
          errors={errors}
        />
      </div>

      {/*<input type='text' /> TODO: GET ALL THE COUNTRIES API AND PREPARE A SELECT FIELD WITH AN AUTOCOMPLETE */}

      <Button
        variant='primary'
        type='submit'
        className='mt-12'
        disabled={isLoading}
      >
        Register
      </Button>
    </Form>
  );
};

export default Register;
