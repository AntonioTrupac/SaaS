import axios from 'axios';
import { useQuery } from 'react-query';

export type APICountryType = {
  name: {
    common: string;
    official: string;
  };
};

export type Country = {
  name?: string;
};

export const fetchAllCountries = async () => {
  const response = await axios.get('https://restcountries.com/v3.1/all', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (response.status >= 400 && response.status <= 500) {
    throw new Error('ERROR! Could not fetch countries!');
  }

  return response.data.map((country: APICountryType) => country.name.common);
};

const useCountries = () => {
  const {
    data,
    isLoading: isLoadingCountries,
    isSuccess: isSuccessCountries,
    isError: isErrorCountries,
    error,
  } = useQuery<Country[]>('countries', fetchAllCountries, {
    refetchOnMount: false,
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
