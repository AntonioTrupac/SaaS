import { yupResolver } from '@hookform/resolvers/yup';
import { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';

import Form from '@/components/form/Form';
import { loginValidationSchema } from '@/components/form/validation';
import Label from '@/components/ui/label/Label';

import { LoginVariables } from '@/graphql';
import Heading from '@/typography/Headings/Heading';

import { State } from './Register';
import Button from '../ui/buttons/Button';
import Input from '../ui/inputs/Input';
import InputError from '../ui/inputs/InputError';
import useLoginMutation from '../../hooks/mutations/useLoginMutation';

export type IFormLoginInput = {
  email: string;
  password: string;
};

const Login = () => {
  const { mutate, isLoading, isError, isSuccess } = useLoginMutation();

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
  } = useForm<IFormLoginInput>({
    defaultValues: {
      email: '',
      password: '',
    },
    mode: 'onTouched',
    reValidateMode: 'onChange',
    resolver: yupResolver(loginValidationSchema),
  });

  const onSubmit: SubmitHandler<IFormLoginInput> = ({ email, password }) => {
    const mappedData: LoginVariables = {
      email,
      password,
    };
    mutate(mappedData);
    reset();
  };
  return (
    <Form
      className='flex w-full flex-col items-center lg:w-[320px]'
      onSubmit={handleSubmit(onSubmit)}
    >
      <Heading variant='h2' className='mb-[18px] text-left'>
        Login
      </Heading>

      <section className='w-full lg:w-[320px]'>
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

        <Button
          variant='primary'
          type='submit'
          className='mt-12 lg:w-[320px]'
          disabled={isLoading}
        >
          Login
        </Button>
      </section>
    </Form>
  );
};

export default Login;
