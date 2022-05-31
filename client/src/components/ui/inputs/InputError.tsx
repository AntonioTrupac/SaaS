import { FieldErrors } from 'react-hook-form';

import { IFormInput } from '@/components/form/Register';
import { InputFieldName } from '@/components/ui/inputs/Input';

type States = {
  isError?: boolean;
  isDisabled?: boolean;
  isSuccess?: boolean;
};

type InputErrorProps = {
  inputFieldName: InputFieldName;
  states: States;
  errors: FieldErrors<IFormInput>;
};

const InputError = ({ states, errors, inputFieldName }: InputErrorProps) => {
  return (
    <>
      {!states.isDisabled && (
        <p
          className={`text-l mt-[8px] text-left text-xs font-light text-red-700 ${
            !errors?.[inputFieldName]?.message ? 'hidden' : 'block'
          }`}
        >
          {errors?.[inputFieldName]?.message}
        </p>
      )}
    </>
  );
};

export default InputError;
