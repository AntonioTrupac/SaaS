import { ComponentPropsWithoutRef } from 'react';

import Paragraph from '@/typography/Paragraph/Paragraph';

type LabelProps = {
  name: string;
  htmlFor: string;
};

type Props = LabelProps & ComponentPropsWithoutRef<'label'>;

const classes: { [p: string]: string | (() => string) } = {};

const Label = ({ name, htmlFor, ...props }: Props) => {
  return (
    <label htmlFor={htmlFor} {...props} className=''>
      <Paragraph size='small' className='mb-[4px] text-left'>
        {name}
      </Paragraph>
    </label>
  );
};

export default Label;
