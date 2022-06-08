import { forwardRef, InputHTMLAttributes, type ReactNode } from 'react';
import { type FieldErrors } from 'react-hook-form';

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
  fieldName: InputFieldName;
  states: States;
  errors: FieldErrors<IFormInput>;
  iconLeft?: ReactNode;
  iconRight?: ReactNode;
};

type Props = InputProps & InputHTMLAttributes<HTMLInputElement>;

const classes = {
  general:
    'flex w-full lg:w-[320px] py-2 px-4 gap-2.5  bg-secondary-100 text-neutrals-900' +
    `placeholder:text-secondary-500 drop-shadow-lg border-none rounded-md focus:outline-none focus:bg-secondary-75 active:bg-secondary-75`,
  focus: (errorMessage: string | undefined) =>
    `focus:outline-none focus:border-primary-200 focus:border focus:ring-2 ${
      errorMessage ? 'focus:ring-danger-400' : 'focus:ring-primary-200'
    }`,
  success:
    'focus:outline-none focus:border-success-400 focus:border focus:ring-2 focus:ring-success-400',
  error: 'ring-2 ring-danger-400 bg-secondary-75',
  disabled: 'disabled:bg-neutrals-60 disabled:shadow-none',
  iconLeft: (iconLeft: ReactNode | undefined) => `${iconLeft ? 'pl-8' : ''}`,
  iconRight: (iconRight: ReactNode | undefined) =>
    `${iconRight ? 'pr-10' : ''}`,
};

const Input = forwardRef<HTMLInputElement, Props>(
  (
    {
      fieldName,
      states: { isError, isDisabled, isSuccess },
      errors,
      iconLeft,
      iconRight,
      ...props
    },
    ref
  ) => {
    const errorMessage = errors?.[fieldName]?.message;

    return (
      <span className='relative block'>
        {iconLeft && (
          <span className='absolute inset-y-0 left-0 flex items-center pl-2'>
            {iconLeft}
          </span>
        )}

        <input
          {...props}
          className={clsxm(
            classes.general,
            classes.focus(errorMessage),
            iconLeft && classes.iconLeft(iconLeft),
            isSuccess && classes.success,
            isError || (errorMessage && classes.error),
            isDisabled && classes.disabled,
            iconRight && classes.iconRight(iconRight),
            'block'
          )}
          ref={ref}
        />

        {iconRight && (
          <span className='absolute inset-y-0 right-4 flex items-center pl-2'>
            {iconRight}
          </span>
        )}
      </span>
    );
  }
);

export default Input;
