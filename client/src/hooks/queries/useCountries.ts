import axios from 'axios';
import { useQuery } from 'react-query';

type Country =
  | {
      id: number;
      name: string;
    }
  | undefined;

const fetchAllCountries = async () => {
  const response = await axios.get('/api/countries');

  if (response.status === 200) {
    return response.data;
  }

  return undefined;
};

const useCountries = () => {
  const {
    data,
    isLoading: isLoadingCountries,
    isSuccess: isSuccessCountries,
    isError: isErrorCountries,
    error,
  } = useQuery<Country[]>('countries', fetchAllCountries, {
    refetchOnMount: true,
    refetchOnWindowFocus: false,
  });

  if (isErrorCountries) {
    console.error(error);
    throw new Error('There has been an error while fetching countries');
  }

  return {
    data,
    isLoadingCountries,
    isSuccessCountries,
    isErrorCountries,
    error,
  };
};

export default useCountries;
