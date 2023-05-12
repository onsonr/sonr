"use client"

import { useToast } from "@/hooks/useToast"

import { AspectRatio } from "@/components/ui/aspect-ratio"
import {
    ContextMenu,
    ContextMenuCheckboxItem,
    ContextMenuContent,
    ContextMenuItem,
    ContextMenuSeparator,
    ContextMenuTrigger,
} from "@/components/ui/context-menu"
import { Icons } from "@/components/icons"
import { QRDialog } from "@/components/explorer/dialog-qr"
import { useEffect, useState } from "react"
import AuroraBackground from "@/components/aurora-background"
import { cn } from "@/lib/utils"
interface CryptoCardProps {
    address: string
    name: string
    type: string,
    className?: string
}
export function CryptoCard({ address, name, type, className }: CryptoCardProps) {
    const [qrOpen, setQROpen] = useState(false)
    const { toast } = useToast()
    return (
        <AspectRatio
            ratio={16 / 9}
            className={cn("bg-clip h-fit w-max rounded-md  bg-white/20 bg-clip-padding shadow-lg backdrop-blur-sm dark:bg-slate-800/40 sm:rounded-xl", className)}
            style={{ position: "relative" }} // add relative positioning
        >
            <QRDialog address={address} open={qrOpen} setOpen={setQROpen} />
            <ContextMenu>
                <ContextMenuTrigger>
                    <div className="rounded-md">

                        <div
                            style={{
                                position: "absolute",
                                top: 0,
                                left: 0,
                                right: 0,
                                bottom: 0,
                                borderRadius: "0.75rem", // add border radius
                            }}
                        >
                            <div className="px-3 py-2">
                                <div className="grid h-fit grid-cols-2 items-stretch justify-between gap-6">
                                    <div className="text-md col-start-3 ml-auto mt-1 pr-1 font-extrabold text-white/60">
                                        {type}
                                    </div>
                                    <div className="col-start-1 pt-4 text-xs font-light text-white/80">
                                        {address.split(":")[2]}
                                    </div>
                                    <div className="text-md col-start-1 pt-6 text-white/60">
                                        {name}
                                    </div>
                                    <div className="col-start-3 pt-6">{getIcon(type)}</div>
                                </div>
                            </div>
                        </div>
                        <AuroraBackground width={350} variant={type.toLowerCase()} />
                    </div>
                </ContextMenuTrigger>
                <ContextMenuContent className="w-64">
                    <ContextMenuItem
                        inset
                        onClick={() => {
                            let t = toast({
                                title: "Copied " + type + " Address",
                                description: address + " address copied to clipboard",
                            })
                            if (!navigator.clipboard) {
                                return
                            }
                            navigator.clipboard.writeText(address)
                            setTimeout(() => {
                                t.dismiss()
                            }, 1000)
                        }}
                    >
                        Copy Address
                    </ContextMenuItem>
                    <ContextMenuItem inset onClick={(e) => {
                        setQROpen(true)
                    }}>
                        Show QR Code
                    </ContextMenuItem>
                    <ContextMenuSeparator />
                    <ContextMenuItem inset disabled>
                        Deposit Funds
                    </ContextMenuItem>
                </ContextMenuContent>
            </ContextMenu>
        </AspectRatio>
    )
}

export function CryptoCardLarge({ address, name, type, className }: CryptoCardProps) {
    const [windowWidth, setWindowWidth] = useState<number>(0);

    useEffect(() => {
        setWindowWidth(window.innerWidth);
        const handleResize = () => setWindowWidth(window.innerWidth);
        window.addEventListener('resize', handleResize);
        return () => window.removeEventListener('resize', handleResize);
    }, []);

    return (
        <div className={cn("bg-clip border-muted-foreground w-full rounded-3xl border sm:w-fit", className)}>
            <div
                style={{
                    position: "absolute",
                    top: 0,
                    left: 0,
                    right: 0,
                    bottom: 0,
                }}
            >
                <div className="h-[500px] items-center p-8 px-6">
                    <div className="grid grid-cols-4 items-stretch justify-between gap-2 md:gap-8 lg:gap-32">
                        <div className="col-start-1 ml-auto mt-1 pr-8 text-base font-extrabold text-white/40 sm:text-xl md:text-2xl lg:text-3xl">
                            {type}
                        </div>
                        <div className="text-white/85 col-start-1 font-mono text-xs font-light sm:text-sm md:text-base lg:text-xl">
                            {address}
                        </div>
                        <div className="col-start-1 mt-12 text-xs text-white/60 sm:text-base md:text-xl lg:text-2xl">
                            {name}
                        </div>
                        <div className="col-start-4 ml-4 mt-8">
                            {getIcon(type, windowWidth < 640 ? 16 : windowWidth < 768 ? 20 : 22)}
                        </div>
                    </div>
                </div>
            </div>
            <AuroraBackground variant={type.toLowerCase()} />
        </div>
    )
}

function getIcon(type: string, size: number = 8) {
    if (type.includes("Ethereum")) {
        return <Icons.ethereum className={cn("text-white/70", `w-${size} h-${size}`)} />
    } else if (type.includes("Bitcoin")) {
        return <Icons.bitcoin className={cn("text-white/70", `w-${size} h-${size}`)} />
    } else if (type.includes("Sonr")) {
        return <Icons.sonr className={cn("text-white/70", `w-${size} h-${size}`)} />
    }
    switch (type) {
        case "ETH":
            return <Icons.ethereum className={cn("text-white/60", `w-${size} h-${size}`)} />
        case "BTC":
            return <Icons.bitcoin className={cn("text-white/60", `w-${size} h-${size}`)} />
        case "SNR":
            return <Icons.sonr className={cn("text-white/60", `w-${size} h-${size}`)} />
        default:
            return <Icons.ethereum className={cn("text-white/60", `w-${size} h-${size}`)} />
    }
}

function getExplorerLink(type: CryptoCardProps["type"], address: string) {
    if (type.includes("Ethereum")) {
        return `https://www.blockchain.com/eth/address/${address}`
    } else if (type.includes("Bitcoin")) {
        return `https://www.blockchain.com/btc/address/${address}`
    } else if (type.includes("Sonr")) {
        return `https://rpc.sonr.ws/address/${address}`
    }
}

export function coinTypeIsBitcoin(type: string) {
    return type.includes("Bitcoin")
}

export function coinTypeIsEthereum(type: string) {
    return type.includes("Ethereum")
}

export function coinTypeIsSonr(type: string) {
    return type.includes("Sonr")
}
