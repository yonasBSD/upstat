import { MonitorItem, columns } from "./Columns";
import { DataTable } from "./DataTable";

export default function DemoPage() {
  const data: MonitorItem[] = [
    {
      name: "Chad",
      url: "chad.com",
      frequency: "30",
      lastStatus: "200",
    },
  ];

  return <DataTable columns={columns} data={data} />;
}
