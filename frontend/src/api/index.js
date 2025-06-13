import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080/api/v1',
});

export const fetchCountryProductStats = (limit=100, offset=0) =>
  api.get(`/country-product-stats?limit=${limit}&offset=${offset}`);

export const fetchRevenueByCountry = () =>
  api.get('/revenue-by-country');

export const fetchTopProducts = (limit=20) =>
  api.get(`/top-products?limit=${limit}`);

export const fetchSalesByMonth = () =>
  api.get('/sales-by-month');

export const fetchRevenueByRegion = (limit=30) =>
  api.get(`/revenue-by-region?limit=${limit}`);
