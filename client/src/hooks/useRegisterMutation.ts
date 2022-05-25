import { useRegister } from '@/graphql';

import { useGraphQLClient } from './useGraphQLClient';

const useRegisterMutation = () => {
  const { graphQLClient } = useGraphQLClient();

  return useRegister(graphQLClient, {
    onMutate: async (register) => {
      console.log(register);
    },
    onSuccess: (data) => {
      console.log(`REGISTRATION SUCCESSFUL!`, data);
    },
    onError: (err) => {
      console.error(err);
    },
  });
};

export default useRegisterMutation;
