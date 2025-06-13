import React from 'react';
import CountryTable from '../components/CountryTable';
import TopProducts from '../components/TopProducts';
import SalesByMonth from '../components/SalesByMonth';
import RegionChart from '../components/RegionChart';

export default function Dashboard() {
  return (
    <div style={{ padding: 20 }}>
      <h1>Global Sales Dashboard</h1>
      <section>
        <h2>Country - Product Revenue</h2>
        <CountryTable />
      </section>
      <section>
        <h2>Top Products</h2>
        <TopProducts />
      </section>
      <section>
        <h2>Monthly Sales Volume</h2>
        <SalesByMonth />
      </section>
      <section>
        <h2>Top Regions by Revenue</h2>
        <RegionChart />
      </section>
    </div>
  );
}
