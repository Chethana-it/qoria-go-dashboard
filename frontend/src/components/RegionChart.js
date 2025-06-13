import React, { useEffect, useState } from 'react';
import { BarChart, Bar, XAxis, YAxis, Tooltip, ResponsiveContainer } from 'recharts';
import { fetchRevenueByRegion } from '../api';
import './RegionChart.css'

export default function RegionChart({ limit=30 }) {
  const [data, setData] = useState([]);
  useEffect(() => {
    fetchRevenueByRegion(limit)
      .then(res => setData(res.data))
      .catch(console.error);
  }, [limit]);

  return (
    <ResponsiveContainer className="region-chart" width="100%" height={300}>
      <BarChart data={data}>
        <XAxis dataKey="Region" tick={{ fontSize: 12 }} />
        <YAxis />
        <Tooltip />
        <Bar dataKey="TotalRevenue" name="Revenue" />
      </BarChart>
    </ResponsiveContainer>
  );
}
