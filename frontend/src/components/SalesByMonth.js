import React, { useEffect, useState } from 'react';
import { LineChart, Line, XAxis, YAxis, Tooltip, ResponsiveContainer } from 'recharts';
import { fetchSalesByMonth } from '../api';
import './SalesByMonth.css'

export default function SalesByMonth() {
  const [data, setData] = useState([]);
  useEffect(() => {
    fetchSalesByMonth()
      .then(res => setData(res.data))
      .catch(console.error);
  }, []);

  return (
    <ResponsiveContainer className="sales-chart" width="100%" height={300}>
      <LineChart data={data}>
        <XAxis dataKey="Month" />
        <YAxis />
        <Tooltip />
        <Line type="monotone" dataKey="QuantitySold" stroke="#8884d8" />
      </LineChart>
    </ResponsiveContainer>
  );
}
