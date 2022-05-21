import * as React from 'react';

import useMoodTypes from '@/hooks/useMoodTypes';

import Layout from '@/components/layout/Layout';
import ButtonLink from '@/components/links/ButtonLink';
import Seo from '@/components/Seo';

export default function HomePage() {
  const { moodTypes } = useMoodTypes();

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
            <div>
              {moodTypes?.map((moodType) => {
                return <div key={moodType.id}>{moodType.name}</div>;
              })}
            </div>
          </div>
        </section>
      </main>
    </Layout>
  );
}
