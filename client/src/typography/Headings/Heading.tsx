import clsxm from '@/lib/clsxm';
import { ComponentPropsWithoutRef, PropsWithChildren } from 'react';

enum HeadingVariant {
  'h1',
  'h2',
  'h3',
}

type HeadingProps = {
  bold?: boolean;
  // keyof typeof infers the type of a js object and returns a type
  // that is the union of its keys
  variant?: keyof typeof HeadingVariant;
  className?: string;
  fontNormal?: boolean;
};

type Props = HeadingProps & ComponentPropsWithoutRef<'h1'>;

const classes = {
  general: 'text-dark',
  h1: (fontNormal: boolean | undefined) =>
    `${
      fontNormal ? 'font-normal' : 'font-bold'
    } text-[64px] tracking-[-0.02em] leading-[75px]`,
  h2: (fontNormal: boolean | undefined) =>
    `${
      fontNormal ? 'font-normal' : 'font-bold'
    } text-[40px] tracking-[-0.02.em] leading-[47px]`,
  h3: (fontNormal: boolean | undefined) =>
    `${
      fontNormal ? 'font-normal' : 'font-bold'
    } text-2xl tracking-[0.02em] leading-6`,
};

const Heading = ({
  children,
  className,
  variant: Heading,
  fontNormal,
}: PropsWithChildren<Props>) => {
  switch (Heading) {
    case 'h1':
    case 'h2':
    case 'h3':
      return (
        <Heading
          className={clsxm(
            className,
            //#region  //*=========== Variants ===========
            [
              Heading === 'h1' && classes.h1(fontNormal),
              Heading === 'h2' && classes.h2(fontNormal),
              Heading === 'h3' && classes.h3(fontNormal),
            ]
            //#endregion  //*======== Variants ===========
          )}
        >
          {children}
        </Heading>
      );
    default:
      return <h1>{children}</h1>;
  }
};

export default Heading;
