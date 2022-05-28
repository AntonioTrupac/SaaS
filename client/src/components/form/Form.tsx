import { ReactNode } from 'react';

type Props = {
  children: ReactNode;
  className?: string;
  onSubmit: () => void;
};

const Form = ({ children, className, onSubmit }: Props) => {
  return (
    <form onSubmit={onSubmit} className={className}>
      {children}
    </form>
  );
};

export default Form;
