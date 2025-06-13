import React, { useEffect, useState } from 'react';
import { fetchCountryProductStats } from '../api';
import './CountryTable.css';

export default function CountryTable() {
  const [data, setData] = useState([]);
  useEffect(() => {
    fetchCountryProductStats(200, 0)
      .then(res => setData(res.data.data))
      .catch(console.error);
  }, []);

  return (
    <div className="table-wrapper">
      <table className="country-table">
        <thead>
          <tr>
            <th>Country</th>
            <th>Product</th>
            <th>Total Revenue</th>
            <th>Transactions</th>
          </tr>
        </thead>
        <tbody>
          {data.map((row, i) => (
            <tr key={i}>
              <td>{row.country}</td>
              <td>{row.product_name}</td>
              <td>{row.total_revenue.toLocaleString()}</td>
              <td>{row.transaction_count}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
