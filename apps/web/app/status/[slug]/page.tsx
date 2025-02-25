import EmptyState from "@/components/empty-state";
import DetailBar from "@/components/status-pages/detail-bar";
import RecentPingChart from "@/components/status-pages/recent-ping-chart";
import { fetchStatusPagesSummary } from "@/lib/api/status-pages";
import { components } from "@/lib/api/types";
import { calculateUptimePercentage, cn } from "@/lib/utils";
import { Activity } from "lucide-react";
import Link from "next/link";
import { redirect } from "next/navigation";

interface PageProps {
  params: Promise<{ slug: string }>;
}

export default async function page({ params }: PageProps) {
  const slug = (await params).slug;

  if (!slug) return redirect("/");

  const data = await fetchStatusPagesSummary(slug);
  if (!data) return null;
  const { monitors, statusPageInfo } = data;

  return (
    <div className="flex flex-col items-center w-full gap-8 p-2 py-8">
      <div className="flex flex-col gap-1 align-center">
        <h1 className="text-2xl font-bold text-center">
          {statusPageInfo?.name}
        </h1>
      </div>

      {!monitors?.length ? (
        <div className="w-full rounded-md max-w-[650px] bg-card">
          <EmptyState
            icon={<Activity />}
            title="No monitors"
            description="The status page has no connected monitors."
          />
        </div>
      ) : (
        <div className="w-full py-4 border rounded-md max-w-[650px] bg-card">
          <div className="w-full text-center">
            <h3 className="text-lg font-medium">Status Check</h3>
          </div>

          <div className="flex flex-col w-full gap-4 mt-8 align-center">
            {monitors.map((m, i) => {
              m?.all?.sort(
                // @ts-ignore
                (a, b) => new Date(b?.timestamp) - new Date(a?.timestamp)
              );

              const uptimePercentage = calculateUptimePercentage(m?.all || []);

              function fillArray(
                arr: (components["schemas"]["HeartbeatSummary"] | undefined)[]
              ) {
                while (arr.length < 45) {
                  arr.push(undefined);
                }
                return arr;
              }
              const finalHeartbeat = fillArray(m?.all || []).reverse();

              return (
                <div
                  className={cn("flex flex-col w-full gap-2", {
                    "border-b pb-4": monitors?.length !== i + 1,
                  })}
                  key={i}
                >
                  <div className="flex justify-between px-4">
                    <span className="font-medium">{m.name}</span>
                    <span
                      className={cn("text-muted-foreground", {
                        "text-green-400": uptimePercentage > 99,
                        "text-yellow-400":
                          95 < uptimePercentage && uptimePercentage < 98,
                        "text-red-400": uptimePercentage <= 95,
                      })}
                    >
                      {uptimePercentage}%
                    </span>
                  </div>
                  <div className="flex flex-col gap-4 px-4">
                    <DetailBar finalHeartbeat={finalHeartbeat} />
                    <RecentPingChart heartbeat={m?.recent || []} />
                  </div>
                </div>
              );
            })}
          </div>
        </div>
      )}

      <span className="text-xs text-muted-foreground">
        Open-source monitoring and status page powered by{" "}
        <Link
          target="_blank"
          href="https://github.com/chamanbravo/upstat"
          className="underline"
        >
          Upstat
        </Link>
        .
      </span>
    </div>
  );
}
