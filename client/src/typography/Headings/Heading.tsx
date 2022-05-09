import { ComponentPropsWithoutRef, PropsWithChildren } from 'react';

import clsxm from '@/lib/clsxm';

enum HeadingVariant {
  'h1',
  'h2',
  'h3',
}

type HeadingProps = {
  isBold?: boolean;
  className?: string;
  // keyof typeof infers the type of a js object and returns a type
  // that is the union of its keys
  variant?: keyof typeof HeadingVariant;
};

type Props = HeadingProps & ComponentPropsWithoutRef<'h1'>;

const classes = {
  general: 'text-dark',
  h1: (isBold: boolean | undefined) =>
    `${
      isBold ? 'font-bold' : 'font-normal'
    } text-[64px] tracking-[-0.02em] leading-[75px]`,
  h2: (isBold: boolean | undefined) =>
    `${
      isBold ? 'font-bold' : 'font-normal'
    } text-[40px] tracking-[-0.02.em] leading-[47px]`,
  h3: (isBold: boolean | undefined) =>
    `${
      isBold ? 'font-bold' : 'font-normal'
    } text-2xl tracking-[0.02em] leading-6`,
};

const Heading = ({
  children,
  className,
  variant: Heading,
  isBold,
}: PropsWithChildren<Props>) => {
  switch (Heading) {
    case 'h1':
    case 'h2':
    case 'h3':
      return (
        <Heading
          className={clsxm(
            className,
            classes.general,
            //#region  //*=========== Variants ===========
            [
              Heading === 'h1' && classes.h1(isBold),
              Heading === 'h2' && classes.h2(isBold),
              Heading === 'h3' && classes.h3(isBold),
            ]
            //#endregion  //*======== Variants ===========
          )}
        >
          {children}
        </Heading>
      );
    default:
      return (
        <h1 className={clsxm(className, classes.h1, classes.general)}>
          {children}
        </h1>
      );
  }
};

export default Heading;
