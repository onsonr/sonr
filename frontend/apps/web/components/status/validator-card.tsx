"use client"

import React from "react";
import {
    Card,
    CardContent,
    CardHeader,
} from "@/components/ui/card";
import {
    Table,
    TableBody,
    TableCaption,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table";
import { SonrNodeResponse } from "@/types/node";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";

interface ValidatorCardProps {
    resp: SonrNodeResponse;
    nickname: string;
}

const ValidatorCard = ({ resp, nickname }: ValidatorCardProps) => {
    // Format the date using JavaScript's Date object
    const blockTime = new Date(resp.result.sync_info.latest_block_time);
    const formattedBlockTime = blockTime.toLocaleString();
    return (
        <Card className="col-span-1">
            <CardHeader>
                <h2 className="font-mono text-3xl font-semibold text-white">{nickname}</h2>
                <div className="mx-32 grid grid-cols-2 gap-x-4 py-2">
                    <Button size="sm" variant={"outline"} onClick={(e)=>{
                        window.open(`https://api.${nickname}.sonr.zone`, "_blank")
                    }}>API</Button>
                    <Button size="sm" variant={"outline"} onClick={(e) => {
                        window.open(`https://rpc.${nickname}.sonr.zone`, "_blank")
                    }}>RPC</Button>
                </div>
            </CardHeader>
            <CardContent>
                <Table>
                    <TableCaption className="text-xs font-extralight text-white/70">{resp.result.node_info.id}</TableCaption>
                    <TableHeader>
                        <TableRow>
                            <TableHead >Latest Block</TableHead>

                            <TableHead>Voting Power</TableHead>
                            <TableHead className="w-[100px]">Synced</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        <TableRow>
                            <TableCell className="text-xs">{formattedBlockTime}</TableCell>
                            <TableCell className="font-mono font-light">{resp.result.validator_info.voting_power}</TableCell>
                            <TableCell className="font-medium"><div className={cn("ml-4 h-4 w-4 rounded-3xl", resp.result.sync_info.catching_up ? "bg-red-500" : "bg-green-500")}></div></TableCell>
                        </TableRow>
                    </TableBody>
                </Table>

            </CardContent>
        </Card>
    )
}


export { ValidatorCard }
