"use client"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog"

import React from "react"
import { DidDocument } from "../../../../packages/client/lib/types"


interface DidDocumentProps {
    did?: DidDocument
}

const DIDProfile = ({ did }: DidDocumentProps) => {
    let username = () => {
        if (did === undefined || did === null) {
            return "anonomous"
        }
        return did.alsoKnownAs ? did.alsoKnownAs[0] : "anonymous"
    }
    return (
        <Card className="max-w-xl">
            <CardHeader>
                <CardTitle className="text-xl font-semibold">{username()}.snr/</CardTitle>
                <div className="grid grid-cols-8 gap-2 pt-2">
                    <Button className="col-span-2 w-full">Send</Button>
                    <Button variant="outline" className="col-span-2">Message</Button>
                </div>
            </CardHeader>
            <CardContent>
                <div className="grid gap-3 md:grid-cols-3 lg:grid-cols-6">
                    <Card className="col-span-2">
                        <CardHeader>
                            <CardTitle>32</CardTitle>
                            <CardDescription>
                                Verifiable Assets
                            </CardDescription>
                        </CardHeader>
                    </Card>
                    <Card className="col-span-2">
                        <CardHeader>
                            <CardTitle>3</CardTitle>
                            <CardDescription>
                                Certificates
                            </CardDescription>
                        </CardHeader>
                    </Card>
                    <Card className="col-span-2">
                        <CardHeader>
                            <CardTitle>1</CardTitle>
                            <CardDescription>
                                Voting History
                            </CardDescription>
                        </CardHeader>
                    </Card>
                </div>
            </CardContent>
        </Card>
    )
}


export { DIDProfile }
