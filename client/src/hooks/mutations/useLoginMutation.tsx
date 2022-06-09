import { useGraphQLClient } from '@/hooks/useGraphQLClient';

import { useLogin } from '@/graphql';

const useLoginMutation = () => {
  const { graphQLClient } = useGraphQLClient();

  return useLogin(graphQLClient, {
    onMutate: async (login) => {
      console.log(login);
    },
    onSuccess: (data) => {
      console.log(`REGISTRATION SUCCESSFUL!`, data);
    },
    onError: (err) => {
      console.error(err);
    },
  });
};

export default useLoginMutation;
