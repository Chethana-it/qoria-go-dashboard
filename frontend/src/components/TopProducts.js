import React, { useEffect, useState } from 'react';
import { BarChart, Bar, XAxis, YAxis, Tooltip, ResponsiveContainer } from 'recharts';
import { fetchTopProducts } from '../api';
import './TopProducts.css';

export default function TopProducts({ limit=20 }) {
  const [data, setData] = useState([]);
  useEffect(() => {
    fetchTopProducts(limit)
      .then(res => setData(res.data))
      .catch(console.error);
  }, [limit]);

  return (
    <ResponsiveContainer className="top-products-chart" width="100%" height={300}>
      <BarChart data={data}>
        <XAxis dataKey="ProductName" tick={{ fontSize: 12 }} />
        <YAxis />
        <Tooltip />
        <Bar dataKey="QuantitySold" name="Quantity Sold" />
      </BarChart>
    </ResponsiveContainer>
  );
}
