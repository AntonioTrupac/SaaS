import { ComponentPropsWithoutRef } from 'react';

import clsxm from '@/lib/clsxm';

import Paragraph from '@/typography/Paragraph/Paragraph';

type LabelProps = {
  name: string;
  htmlFor: string;
  className?: string;
};

type Props = LabelProps & ComponentPropsWithoutRef<'label'>;

const Label = ({ name, htmlFor, className, ...props }: Props) => {
  return (
    <label htmlFor={htmlFor} {...props} className=''>
      <Paragraph
        size='normal'
        className={clsxm('mb-[4px] text-left', className)}
      >
        {name}
      </Paragraph>
    </label>
  );
};

export default Label;
