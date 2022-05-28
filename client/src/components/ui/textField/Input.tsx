import { type ComponentPropsWithoutRef } from 'react';
import { type FieldErrors, UseFormRegister } from 'react-hook-form';

import clsxm from '@/lib/clsxm';

import { type IFormInput } from '@/components/form/Register';

type States = {
  isError?: boolean;
  isDisabled?: boolean;
  isSuccess?: boolean;
};
export type InputType = 'text' | 'email' | 'password';
export type InputFieldName =
  | 'email'
  | 'password'
  | 'firstName'
  | 'lastName'
  | 'country'
  | 'city';

type InputProps = {
  type?: InputType;
  register: UseFormRegister<IFormInput>;
  inputFieldName: InputFieldName;
  states: States;
  errors: FieldErrors<IFormInput>;
};

type Props = InputProps & ComponentPropsWithoutRef<'input'>;

const classes = {
  general:
    'flex w-full lg:w-[320px] py-2 px-4 gap-2.5  bg-secondary-100 text-neutrals-900 ' +
    `placeholder:text-secondary-500 shadow-md border-none rounded-md focus:outline-none focus:bg-secondary-75 active:bg-secondary-75`,
  focus: (errorMessage: string | undefined) =>
    `focus:outline-none focus:border-primary-200 focus:border focus:ring-2 ${
      errorMessage ? 'focus:ring-danger-400' : 'focus:ring-primary-200'
    }`,
  success:
    'focus:outline-none focus:border-success-400 focus:border focus:ring-2 focus:ring-success-400',
  error: 'ring-2 ring-danger-400 bg-secondary-75',
  disabled: 'disabled:bg-neutrals-60 disabled:shadow-none',
};

const Input = ({
  placeholder,
  type = 'text',
  inputFieldName,
  register,
  states: { isError, isDisabled, isSuccess },
  errors,
  ...props
}: Props) => {
  const errorMessage = errors?.[inputFieldName]?.message;
  return (
    <>
      <input
        {...props}
        placeholder={placeholder}
        disabled={isDisabled}
        type={type}
        {...register(inputFieldName)}
        className={clsxm(
          classes.general,
          classes.focus(errorMessage),
          isSuccess && classes.success,
          isError || (errorMessage && classes.error),
          isDisabled && classes.disabled
        )}
      />
    </>
  );
};

export default Input;
