import { type PropsWithChildren } from 'react';
import { FieldErrors, UseFormRegister } from 'react-hook-form';

import { type IFormInput } from '@/ui/form/Register';
import Input, {
  type InputFieldName,
  type InputType,
} from '@/ui/textField/Input';
import Label from '@/ui/textField/Label';

type States = {
  isError?: boolean;
  isDisabled?: boolean;
  isSuccess?: boolean;
};

type TextFieldProps = {
  type?: InputType;
  register: UseFormRegister<IFormInput>;
  inputFieldName: InputFieldName;
  placeholder: string;
  states: States;
  errors: FieldErrors<IFormInput>;
};

const TextField = ({
  placeholder,
  type,
  register,
  inputFieldName,
  states,
  errors,
}: PropsWithChildren<TextFieldProps>) => {
  return (
    <div className='my-4'>
      <Label htmlFor={inputFieldName} name={placeholder} />

      <Input
        placeholder={placeholder}
        type={type}
        register={register}
        inputFieldName={inputFieldName}
        states={states}
        errors={errors}
      />

      {!states.isDisabled && (
        <p className='text-l mt-[8px] text-left text-xs font-light text-red-700'>
          {errors?.[inputFieldName]?.message}
        </p>
      )}
    </div>
  );
};

export default TextField;
