import { Metadata } from "next"

import { Search } from "@/components/explorer/search"
import AccountSwitcher from "@/components/dashboard/account-switcher"
import { UserNav } from "@/components/dashboard/user-nav"
import { DashboardNav } from "@/components/main-nav"
import { useSonr } from "@/hooks/useSonr"
import { DIDProfile } from "@/components/explorer/did-profile"
import { useEffect, useState } from "react"
import { DidDocument } from "../../../../packages/client/lib/types"
import { SonrClient } from "../../../../packages/client/lib"
import useSWR, { SWRConfig, preload } from 'swr'


const fetcher = (url) => fetch(url).then((res) => res.json())

export const metadata: Metadata = {
    title: "Search",
    description: "Example dashboard app using the components.",
}

export async function getStaticProps() {
    // `getStaticProps` is executed on the server side.
    let origin = "sonr.id"
    if (process.env.NODE_ENV === "development") {
        origin = "localhost"
    }
    const client = new SonrClient(origin);
    const article = await client.did.list();
    return {
        props: {
            fallback: {
                '/api/did': article
            }
        }
    }
}


export default function SearchPage() {
    const sonr = useSonr()
    const [didList, setDidList] = useState<DidDocument[] | null>(null);
    const { data } = useSWR('/api/article', fetcher)

    return (
        <SWRConfig value={{  }}>
            <div className="hidden flex-col md:flex">
                <div className="border-b dark:border-slate-700">
                    <div className="flex h-16 items-center px-4">
                        <div className="mr-6">
                            <AccountSwitcher sonr={sonr} />
                        </div>
                        <DashboardNav className="mx-6" />
                        <div className="ml-auto flex items-center space-x-4">
                            <Search />
                            <UserNav />
                        </div>
                    </div>
                    <section className="p-24">
                        {
                            didList && didList.map((did) => {
                                return (
                                    <div key={did.id}>
                                        <DIDProfile did={did} />
                                    </div>
                                )
                            })
                        }
                    </section>
                </div>
            </div>
        </SWRConfig>
    )
}
