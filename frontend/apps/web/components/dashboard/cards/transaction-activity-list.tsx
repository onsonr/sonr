"use client"
import { ChevronDown } from "lucide-react"

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import {
    Command,
    CommandEmpty,
    CommandGroup,
    CommandInput,
    CommandItem,
    CommandList,
} from "@/components/ui/command"
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from "@/components/ui/popover"
import { Icons } from "@/components/icons"
import { Separator } from "@/components/ui/separator"

export function TransactionActivityList() {
    return (
        <Card className="col-span-6">
            <CardHeader>
                <CardTitle>Transaction Activity</CardTitle>
                <CardDescription>
                    Invite your team members to collaborate.
                </CardDescription>
            </CardHeader>
            <CardContent className="grid gap-6">
                <TransactionActivityListItem from="Jane Cooper" to="jc@email.com" type="deposit" amount="200.000" />

            </CardContent>
        </Card>
    )
}

interface TransactionActivityListItemProps {
    from: string
    to: string
    amount: string
    type: 'deposit' | 'withdrawal' | 'swap' | 'transfer'
}

export function TransactionActivityListItem({ from, to, amount, type }: TransactionActivityListItemProps) {

    const FallbackIcon = () => {
        switch (type) {
            case 'deposit':
                return <AvatarFallback className="bg-green-100 dark:bg-green-700"> <Icons.deposit className="h-4 w-4" /> </AvatarFallback>
            case 'withdrawal':
                return <AvatarFallback className="bg-red-100 dark:bg-red-700"> <Icons.withdrawal className="h-4 w-4" /> </AvatarFallback >
            case 'swap':
                return <AvatarFallback className="bg-green-100 dark:bg-green-700"> <Icons.swap className="h-4 w-4" /> </AvatarFallback >
            case 'transfer':
                return <AvatarFallback className="bg-red-100 dark:bg-red-700"> <Icons.send className="h-4 w-4" /> </AvatarFallback >
        }
    }

    const TitleMessage = () => {
        switch (type) {
            case 'deposit':
                return <p className="text-sm font-medium leading-none">{`Deposited ${amount} from ${from} into ${to}`}</p>
            case 'withdrawal':
                return <p className="text-sm font-medium leading-none">{`Withdrew ${amount} from ${from} into ${to}`}</p>
            case 'swap':
                return <p className="text-sm font-medium leading-none">{`Swapped ${amount} from ${from} into ${to}`}</p>
            case 'transfer':
                return <p className="text-sm font-medium leading-none">{`Sent ${amount} from ${from} into ${to}`}</p>
        }
    }

    const SubtitleText = () => {
        switch (type) {
            case 'deposit':
                return <p className="text-muted-foreground pt-1 text-xs">{`+ ${amount} SNR`}</p>
            case 'withdrawal':
                return <p className="text-x text-muted-foreground pt-1">{`- ${amount} SNR`}</p>
            case 'swap':
                return <p className="text-muted-foreground pt-1 text-xs">{`+ ${amount} SNR`}</p>
            case 'transfer':
                return <p className="text-muted-foreground pt-1 text-xs">{`- ${amount} SNR`}</p>
        }
    }


    return (
        <div className="pt-2">
            <div className="flex items-center justify-between space-x-4">
                <div className="flex items-center space-x-4">
                    <Avatar>
                        <AvatarImage src="/avatars/02.png" />
                        <AvatarFallback><FallbackIcon /></AvatarFallback>
                    </Avatar>
                    <div>
                        <TitleMessage />
                        <SubtitleText />
                    </div>
                </div>
                <Button variant={"subtle"}>
                    View
                </Button>


            </div>
            <Separator className="mt-3" />
        </div>

    )
}
