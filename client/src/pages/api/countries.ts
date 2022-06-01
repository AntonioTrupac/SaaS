import axios from 'axios';
import type { NextApiRequest, NextApiResponse } from 'next';

type Country = {
  name: {
    common: string;
    official: string;
  };
};

type Data = {
  country?: string[];
  message: string;
};

const fetchCountries = async () => {
  const response = await axios.get('https://restcountries.com/v3.1/all', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (response.status >= 400 && response.status <= 500) {
    throw new Error('ERROR! Could not fetch countries!');
  }
  return response.data.map((country: Country) => country.name.official);
};

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  try {
    const data = await fetchCountries();
    console.log('IN TRY', data);
    res.status(200).json({
      country: data,
      message: 'Response successful',
    });
  } catch (err) {
    res.status(404).json({ country: undefined, message: 'Response failed' });
    throw new Error('Error! No countries found!');
  }
}
