import { ComponentPropsWithRef, PropsWithChildren } from 'react';

import clsxm from '@/lib/clsxm';

enum Size {
  'small',
  'normal',
}

type ParagraphProps = {
  size?: keyof typeof Size;
  isBold?: boolean;
  className?: string;
};

type Props = ParagraphProps & ComponentPropsWithRef<'p'>;

const classes = {
  small: 'text-sm leading-[22px]',
  normal: 'text-base leading-[22px]',
};

const Paragraph = ({
  children,
  size = 'normal',
  isBold = false,
  className,
}: PropsWithChildren<Props>) => {
  return (
    <p
      className={clsxm(
        className,
        isBold && 'font-bold',
        //#region  //*=========== sizes ===========
        [size === 'small' && classes.small, size === 'normal' && classes.normal]
        //#endregion  //*======== sizes ===========
      )}
    >
      {children}
    </p>
  );
};

export default Paragraph;
