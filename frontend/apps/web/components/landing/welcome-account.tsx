"use client"

import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"

import { CryptoCard } from "@/components/card-crypto"
import { DidDocument } from "../../../../packages/client/lib/types"
import { Organization } from "@/types/org"

export interface WelcomeAccountProps {
    did: DidDocument
    org?: Organization
    existing: boolean
}

export function WelcomeAccount({ did, org, existing = false }: WelcomeAccountProps) {
    return (
        <Card className="max-w-2xl">
            <CardHeader className="justify-center align-middle">
                <div className="lg:ml-6">
                    <CryptoCard address={did.id} type="Sonr" name={did.alsoKnownAs ? did.alsoKnownAs[0] : ""} />
                </div>

            </CardHeader>
            <CardContent className="grid gap-4">
                {
                    existing ? (
                        <CardTitle className="text-2xl">Hey there! {did.alsoKnownAs ? did.alsoKnownAs[0] : ""}</CardTitle>
                    ) : (
                            <CardTitle className="text-2xl">Welcome, {did.alsoKnownAs ? did.alsoKnownAs[0] : ""}</CardTitle>
                    )
                }

                {org != null && org.amount > 0 ? (
                    <CardDescription>
                        You have successfully created your Sonr account on behalf of, <span className="font-semibold">{org.name}</span>.
                        <span>
                            {" "} You have been awarded <span className="font-mono">{org.amount} SNR</span> for creating your account.
                        </span>
                    </CardDescription>
                ) : (
                    <CardDescription>
                        Welcome back to your Sonr account. You can now use your Sonr account to sign in to any Sonr enabled application.
                    </CardDescription>
                )
                }
                {org != null ?? (
                    <CardDescription>
                        You can now use your Sonr account to sign in to any Sonr enabled application.
                    </CardDescription>
                )}

                <div className="grid gap-2">
                </div>
                <div className="grid gap-2">
                    <div className="space-y-1">
                    </div>
                </div>
            </CardContent>
            <CardFooter className="w-full justify-center">
                <Button className="w-full" onClick={() => {
                    window.location.href = "/wallet"
                }}>
                    Continue to Dashboard
                </Button>
            </CardFooter>
        </Card>
    )
}
