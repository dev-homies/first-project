import axios from 'axios';

export const Axios = axios.create({
  baseURL: `${process.env.NEXT_PUBLIC_API_HOST}/v1`,
  timeout: 10000,
  withCredentials: true,
});
