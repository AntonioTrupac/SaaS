import { PropsWithChildren } from 'react';
import { UseFormRegister } from 'react-hook-form';

import { IFormInput } from '@/ui/form/Register';
import Input, {
  type InputFieldName,
  type InputType,
  type InputVariant,
} from '@/ui/textField/Input';
import Label from '@/ui/textField/Label';

type TextFieldProps = {
  variant?: keyof typeof InputVariant;
  type?: InputType;
  register: UseFormRegister<IFormInput>;
  inputFieldName: InputFieldName;
  disabled: boolean;
  placeholder: string;
};

const TextField = ({
  placeholder,
  disabled = false,
  type,
  register,
  inputFieldName,
}: PropsWithChildren<TextFieldProps>) => {
  return (
    <div className='my-4'>
      <Label htmlFor={inputFieldName} name={placeholder} />
      <Input
        placeholder={placeholder}
        disabled={disabled}
        type={type}
        register={register}
        inputFieldName={inputFieldName}
      />
    </div>
  );
};

export default TextField;
