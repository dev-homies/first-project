import axios from 'axios';

export const Axios = axios.create({
  baseURL: 'http://localhost:4000/v1',
  timeout: 10000,
  withCredentials: true,
});
