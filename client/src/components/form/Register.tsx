import { yupResolver } from '@hookform/resolvers/yup';
import { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';

import useRegisterMutation from '@/hooks/mutations/useRegisterMutation';

import { registerValidationSchema } from '@/components/form/validation';
import Button from '@/components/ui/buttons/Button';
import Input from '@/components/ui/inputs/Input';
import InputError from '@/components/ui/inputs/InputError';
import Label from '@/components/ui/label/Label';

import { RegisterVariables } from '@/graphql';
import Heading from '@/typography/Headings/Heading';

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
    mode: 'onTouched',
    reValidateMode: 'onChange',
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
      <Heading variant='h2' isBold className='mb-[18px] text-left'>
        Sign in
      </Heading>

      <div className='my-1.5 w-full'>
        <Label htmlFor='firstName' name='First Name' />
        <Input
          type='text'
          {...register('firstName')}
          fieldName='firstName'
          placeholder='First Name'
          states={mutationStates}
          errors={errors}
        />
        <InputError
          errors={errors}
          inputFieldName='firstName'
          states={mutationStates}
        />
      </div>

      <div className='my-1.5 w-full'>
        <Label htmlFor='lastName' name='Last Name' />
        <Input
          type='text'
          {...register('lastName')}
          fieldName='lastName'
          placeholder='Last Name'
          states={mutationStates}
          errors={errors}
        />
        <InputError
          errors={errors}
          inputFieldName='lastName'
          states={mutationStates}
        />
      </div>

      <div className='my-1.5 w-full'>
        <Label htmlFor='email' name='Email' />
        <Input
          type='email'
          {...register('email')}
          fieldName='email'
          placeholder='Email'
          states={mutationStates}
          errors={errors}
        />
        <InputError
          errors={errors}
          inputFieldName='email'
          states={mutationStates}
        />
      </div>

      <div className='my-1.5 w-full'>
        <Label htmlFor='password' name='Password' />
        <Input
          type='password'
          {...register('password')}
          fieldName='password'
          placeholder='Password'
          states={mutationStates}
          errors={errors}
        />
        <InputError
          errors={errors}
          inputFieldName='password'
          states={mutationStates}
        />
      </div>

      <div className='my-1.5 w-full'>
        <Label htmlFor='country' name='Country' />
        <Input
          type='text'
          {...register('country')}
          fieldName='country'
          placeholder='Country'
          states={mutationStates}
          errors={errors}
        />
        <InputError
          errors={errors}
          inputFieldName='country'
          states={mutationStates}
        />
      </div>

      <div className='mt-1.5 w-full'>
        <Label htmlFor='city' name='City' />
        <Input
          type='text'
          {...register('city')}
          fieldName='city'
          placeholder='City'
          states={mutationStates}
          errors={errors}
        />
        <InputError
          errors={errors}
          inputFieldName='city'
          states={mutationStates}
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
