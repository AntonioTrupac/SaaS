import { ComponentPropsWithoutRef } from 'react';
import { UseFormRegister } from 'react-hook-form';

import clsxm from '@/lib/clsxm';

import { IFormInput } from '@/ui/form/Register';

export enum InputVariant {
  'default',
  'active',
  'success',
  'danger',
  'disabled',
}

export type InputType = 'text' | 'email' | 'password';
export type InputState = 'success' | 'error' | 'disabled';
export type InputFieldName =
  | 'email'
  | 'password'
  | 'firstName'
  | 'lastName'
  | 'country'
  | 'city';

type InputProps = {
  variant?: keyof typeof InputVariant;
  type?: InputType;
  register: UseFormRegister<IFormInput>;
  inputFieldName: InputFieldName;
};

type Props = InputProps & ComponentPropsWithoutRef<'input'>;

const classes: { [p: string]: string | (() => string) } = {
  general:
    'flex w-full lg:w-[320px] py-2 px-4 gap-2.5 bg-secondary-100 text-neutrals-900 ' +
    'placeholder:text-secondary-500 shadow-md border-none rounded-md focus:outline-none',
  active:
    'focus:outline-none focus:border-primary-200 focus:border focus:ring-2 focus:ring-primary-200',
};

const Input = ({
  disabled,
  placeholder,
  type = 'text',
  inputFieldName,
  register,
  ...props
}: Props) => {
  return (
    <input
      {...props}
      placeholder={placeholder}
      disabled={disabled}
      type={type}
      {...register(inputFieldName)}
      className={clsxm(classes.general, classes.active)}
    />
  );
};

export default Input;
