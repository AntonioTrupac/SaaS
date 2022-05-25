import * as React from 'react';

import Layout from '@/components/layout/Layout';
import Seo from '@/components/Seo';

import Register from '@/ui/form/Register';
import ButtonLink from '@/ui/links/ButtonLink';

export default function HomePage() {
  return (
    <Layout>
      {/* <Seo templateTitle='Home' /> */}
      <Seo />

      <main>
        <section className='bg-secondary-50'>
          <div className='layout flex min-h-screen flex-col items-center justify-center text-center'>
            <ButtonLink className='mt-6' href='/components' variant='light'>
              See all components
            </ButtonLink>

            <Register />
          </div>
        </section>
      </main>
    </Layout>
  );
}
