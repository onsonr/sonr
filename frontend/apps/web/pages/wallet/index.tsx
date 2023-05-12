
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Search } from "@/components/explorer/search"
import AccountSwitcher from "@/components/dashboard/account-switcher"
import { UserNav } from "@/components/dashboard/user-nav"
import { DashboardNav } from "@/components/main-nav"
import GridOverview from "@/components/dashboard/tabs/grid-overview"
import { useSonr } from "@/hooks/useSonr"
import GridTransactions from "@/components/dashboard/tabs/grid-transactions"
import GridInbox from "@/components/dashboard/tabs/grid-inbox"
import { Icons } from "@/components/icons"
import { Logo } from "@/components/logo"
import { useState } from "react"
import { CommandMenu } from "@/components/dashboard/command-menu"
import { Footer } from "@/components/footer"
import va from '@vercel/analytics';

export default function WalletPage({ data }) {
    const sonr = useSonr()
    const [cmdIsOpen, setCmdIsOpen] = useState<boolean>(false);
    return (
        <>
            <CommandMenu open={cmdIsOpen} setOpen={setCmdIsOpen} />
            <div className="hidden flex-col md:flex">
                <div className="border-b py-1.5 dark:border-gray-700/40 dark:bg-black/60 md:mx-[8%] lg:mx-[20%]">
                    <div className="flex h-16 items-center px-4">
                        <div className="mr-[2%]">
                            <Logo />
                        </div>
                        <DashboardNav className="mx-6" />
                        <div className="ml-auto flex items-center space-x-4">
                            <Search />
                            <UserNav />
                        </div>
                    </div>
                </div>
                <div className="flex-1 space-y-4 self-center p-8 pt-24">
                    <div className="flex items-center justify-between space-y-2">
                        <div className="p-1">
                            <h2 className="text-4xl font-bold tracking-tight opacity-80">Welcome, {sonr.user?.username}</h2>
                            <p className="ml-1 mt-1 font-mono text-sm text-slate-600 dark:text-slate-400">{sonr.user?.did}</p>
                        </div>
                        <div className="flex items-center space-x-2">
                            <AccountSwitcher sonr={sonr} />
                        </div>
                    </div>
                    <Tabs defaultValue="overview" className="space-y-4 md:w-[800px] lg:w-[950px]" onValueChange={(v) => {
                        va.track('WalletDashboardChange', { value: v });
                    }}>
                        <TabsList className="mb-8">
                            <TabsTrigger value="overview">
                                <Icons.dashboard className="mr-1.5 h-4" />
                                Overview
                            </TabsTrigger>
                            <TabsTrigger value="inbox">
                                <Icons.inbox className="mr-1.5 h-4" />
                                Inbox
                            </TabsTrigger>
                            <TabsTrigger value="transactions">
                                <Icons.activity className="mr-1.5 h-4" />
                                Payments
                            </TabsTrigger>
                        </TabsList>
                        <TabsContent value="overview" className="space-y-4">
                            <GridOverview />
                        </TabsContent>
                        <TabsContent value="inbox" className="space-y-4">
                            <GridInbox />
                        </TabsContent>
                        <TabsContent value="transactions" className="space-y-4">
                            <GridTransactions />
                        </TabsContent>
                    </Tabs>
                </div>
            </div>
            <Footer />
        </>
    )
}
