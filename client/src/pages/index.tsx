import * as React from 'react';

import Register from '@/components/form/Register';
import Layout from '@/components/layout/Layout';
import Seo from '@/components/Seo';

export default function HomePage() {
  return (
    <Layout>
      {/* <Seo templateTitle='Home' /> */}
      <Seo />

      <main>
        <section className='bg-secondary-50'>
          <div className='layout flex min-h-screen flex-col items-center justify-center text-center'>
            <Register />
          </div>
        </section>
      </main>
    </Layout>
  );
}
