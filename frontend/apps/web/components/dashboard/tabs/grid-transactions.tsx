"use client"

import { DemoReportAnIssue } from "@/components/dashboard/cards/report-an-issue";
import { TransactionActivityList } from "@/components/dashboard/cards/transaction-activity-list";

export default function GridTransactions() {
    return (
        <>
            <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-10">
                <TransactionActivityList />
                <DemoReportAnIssue />
            </div>
        </>
    )
}
