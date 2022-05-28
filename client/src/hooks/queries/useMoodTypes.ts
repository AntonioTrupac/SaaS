import { useGraphQLClient } from '@/hooks/useGraphQLClient';
import useQueryStatusLogging from '@/hooks/useQueryStatusLogging';

import { useGetMoodTypes } from '@/graphql';

const useMoodTypes = () => {
  const { graphQLClient } = useGraphQLClient();

  const queryInfo = useGetMoodTypes(
    graphQLClient,
    {},
    {
      useErrorBoundary: true,
      onSuccess: () => {
        // eslint-disable-next-line no-console
        console.log(Date.now(), `Fetching mood types succeeded`);
      },
    }
  );

  useQueryStatusLogging(queryInfo, `Mood type details`);

  const moodTypes = queryInfo.data?.getMoodTypes;

  return {
    ...queryInfo,
    moodTypes,
  };
};

export default useMoodTypes;
