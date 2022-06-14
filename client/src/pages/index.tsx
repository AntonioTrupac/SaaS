import { GetStaticProps } from 'next';
import * as React from 'react';
import { dehydrate, DehydratedState, QueryClient } from 'react-query';

import { fetchAllCountries } from '@/hooks/queries/useCountries';

import { Login } from '@/components/form';
import Layout from '@/components/layout/Layout';
import Seo from '@/components/Seo';

export const getStaticProps: GetStaticProps = async (): Promise<{
  props: { dehydratedState: DehydratedState };
}> => {
  const queryClient = new QueryClient();

  await queryClient.prefetchQuery('countries', fetchAllCountries);
  return {
    props: {
      dehydratedState: dehydrate(queryClient),
    },
  };
};

export default function HomePage() {
  return (
    <Layout>
      <Seo />
      <main>
        <section className='bg-secondary-50'>
          <div className='layout flex min-h-screen flex-col items-center justify-center text-center'>
            {/*<Register />*/}
            <Login />
          </div>
        </section>
      </main>
    </Layout>
  );
}
