import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';

import Form from '@/components/form/Form';
import { loginValidationSchema } from '@/components/form/validation';
import Label from '@/components/ui/label/Label';

import Heading from '@/typography/Headings/Heading';

export type IFormLoginInput = {
  email: string;
  password: string;
};

const Login = () => {
  // const [mutationStates] = useState<State>({
  //   isLoading,
  //   isError,
  //   isSuccess,
  // });

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
    reset();
  };
  return (
    <Form onSubmit={handleSubmit(onSubmit)}>
      <Heading variant='h2' className='mb-[18px] text-left'>
        Login
      </Heading>

      <section className='w-full lg:w-[320px]'>
        <div className='my-1.5 w-full'>
          <Label htmlFor='email' name='Email' />
          {/*<Input fieldName={} states={} errors={} />*/}
          {/*<InputError*/}
          {/*  inputFieldName='email'*/}
          {/*  states={mutationStates}*/}
          {/*  errors={}*/}
          {/*/>*/}
        </div>
      </section>
    </Form>
  );
};

export default Login;
