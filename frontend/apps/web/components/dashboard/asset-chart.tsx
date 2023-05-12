"use client"

import { Button } from "@/components/ui/button";
import { Bar, BarChart, Cell, Legend, Pie, PieChart, ResponsiveContainer, XAxis, YAxis } from "recharts"

const data = [
  {
    name: "SNR",
    value: 5000,
  },
  {
    name: "BTC",
    value: 3600,
  },
  {
    name: "ETH",
    value: 4320,
  },
  {
    name: "USDC",
    value: 1200,
  }
]

const COLORS = [
  "#2BC0F6",
  "#D54E18",
  "#6D9095",
  "#51CEA9",
]

const RADIAN = Math.PI / 180;
const renderCustomizedLabel = ({ cx, cy, midAngle, innerRadius, outerRadius, percent, index }) => {
  const radius = innerRadius + (outerRadius - innerRadius) * 0.5;
  const x = cx + radius * Math.cos(-midAngle * RADIAN);
  const y = cy + radius * Math.sin(-midAngle * RADIAN);

  return (
    <>
      <text x={x} y={y} fill="white" textAnchor={x > cx ? 'start' : 'end'} dominantBaseline="central" className="font-mono text-xs">
        {`${(percent * 100).toFixed(0)}%`}
      </text>
    </>
  );
};

export function AssetPieChart() {
  return (
    <ResponsiveContainer width="100%" height={300}>
      <PieChart>
        <Legend align="center" verticalAlign="bottom" height={36} />
        <Pie
          data={data}
          label={renderCustomizedLabel}
          dataKey="value"
          nameKey="name"
          labelLine={false}
          cx="50%"
          cy="50%"
          startAngle={0}
          endAngle={360}
          outerRadius={85}
        >
          {data.map((entry, index) => (
            <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
          ))}
        </Pie>
      </PieChart>
    </ResponsiveContainer>
  )
}
