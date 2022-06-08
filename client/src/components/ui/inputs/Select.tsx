import { ComponentPropsWithRef, forwardRef } from 'react';
import { type UseFormRegister } from 'react-hook-form';

import clsxm from '@/lib/clsxm';
import { Country } from '@/hooks/queries/useCountries';

// import { Country } from '@/hooks/queries/useCountries';
import { type IFormInput } from '@/components/form/Register';

type States = {
  isError?: boolean;
  isDisabled?: boolean;
  isSuccess?: boolean;
};

type SelectProps = {
  country?: Country[];
  states: States;
  fieldName?: string;
  errorMessage?: string;
};

type Props = SelectProps &
  ReturnType<UseFormRegister<IFormInput>> &
  ComponentPropsWithRef<'select'>;

const classes = {
  general:
    'flex w-full lg:w-[320px] py-2 pl-4 pr-8 gap-2.5  bg-secondary-100 text-neutrals-900' +
    `placeholder:text-secondary-500 drop-shadow-lg border-none rounded-md focus:outline-none focus:bg-secondary-75 active:bg-secondary-75`,
  focus: (errorMessage: string | undefined) =>
    `focus:outline-none focus:border-primary-200 focus:border focus:ring-2 ${
      errorMessage ? 'focus:ring-danger-400' : 'focus:ring-primary-200'
    }`,
  error: 'ring-2 ring-danger-400 bg-secondary-75',
};

const Select = forwardRef<HTMLSelectElement, Props>(
  (
    {
      states: { isError, isDisabled, isSuccess },
      errorMessage,
      onChange,
      onBlur,
      name,
      country,
    },
    ref
  ) => {
    return (
      <div>
        <select
          className={clsxm(
            classes.general,
            classes.focus(errorMessage),
            isError || (errorMessage && classes.error)
          )}
          name={name}
          ref={ref}
          onChange={onChange}
          onBlur={onBlur}
        >
          <option value='' disabled hidden>
            Choose here
          </option>
          {country?.map((country, index) => {
            return (
              <option
                className='truncate'
                key={`country-${index}`}
                value={country.name}
              >
                {country}
              </option>
            );
          })}
        </select>
      </div>
    );
  }
);

export default Select;
