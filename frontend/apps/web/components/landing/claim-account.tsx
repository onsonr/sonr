"use client"
import { get } from '@vercel/edge-config';
import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
    StyledCard,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Icons } from "@/components/icons"
import { useState } from "react"
import { useSonr } from "@/hooks/useSonr"

import { Separator } from "@/components/ui/separator"
import Link from "next/link"
import { DidDocument } from "../../../../packages/client/lib/types"
import { Organization } from '@/types/org';
import va from "@vercel/analytics"

export interface ClaimAccountProps {
    handleRegisterComplete: (did: string, didDocument: DidDocument) => Promise<void>;
    handleLoginComplete: (did: string, didDocument: DidDocument) => Promise<void>;
    handleCredentialSet: (credential: PublicKeyCredential) => Promise<void>;
    org: Organization | undefined;
    setOrg: (org: Organization) => void;
}

export function ClaimAccount({ handleRegisterComplete, handleCredentialSet, org, setOrg, handleLoginComplete }: ClaimAccountProps) {
    const [isAvailable, setIsAvailable] = useState<boolean | null>(null);
    const [alias, setAlias] = useState("");
    const [code, setCode] = useState("");
    const [hasValidCode, setHasValidCode] = useState<boolean | null>(null);
    const [isSignInVisible, setIsSignInVisible] = useState<boolean>(false);

    const [isProspective, setIsProspective] = useState<boolean>(false);
    const [prospectName, setProspectName] = useState<string>("");
    const [prospectCompany, setProspectCompany] = useState<string>("");


    const sonr = useSonr()
    const checkAvailability = async (v: string) => {
        va.track("check-alias", { "alias": v });
        const resp = await sonr.checkAlias(v);
        setIsAvailable(resp.available && v.length >= 4);
    };

    const onRegister = async () => {
        va.track("claim-account", { "alias": alias, "code": code, "prospectName": prospectName, "prospectCompany": prospectCompany, "isProspective": isProspective, "hasValidCode": hasValidCode, "isAvailable": isAvailable });
        await sonr.register({ alias, onCredentialSet: handleCredentialSet, onRegisterComplete: handleRegisterComplete });
    };

    const onAccessCode = async () => {
        let resp = await fetch(`/api/auth/code?code=${code}`, {
            method: "GET",
        });
        let data = await resp.json();
        setHasValidCode(data.success);
        let org = data.org;
        if (data.success) {
            console.log(data);
            setOrg(org);
        }
        // setIsProspective(org.name === "Prospective Investor");
        va.track("check-code", { "code": code, "success": data.success, "org": org });
    }

    const checkDisabled = () => {
        return alias.length < 4 || isAvailable === false || !hasValidCode || (isProspective && (prospectName.length === 0 || prospectCompany.length === 0))
    }

    return (
        <>
            {!isSignInVisible && (
                <Card className="max-w-3xl pb-4 shadow-sm dark:border-slate-600/40 dark:bg-black/10 dark:shadow-lg">
                    <CardContent className="grid gap-4 gap-y-6 pt-8">
                        <div className="grid gap-2">
                            <Label htmlFor="password">Invite Code</Label>
                            <div className="relative">
                                <div className='grid grid-cols-6 gap-x-2'>
                                    <Input disabled={hasValidCode != null && hasValidCode} className='col-span-5' id="password" placeholder="AX13fA" required maxLength={6} onChange={(v) => {
                                        setCode(v.target.value.toUpperCase())
                                    }} onFocus={onAccessCode} />
                                    <Button disabled={hasValidCode != null && hasValidCode} className='col-span-1' variant="ghost" onClick={onAccessCode}>
                                        Check
                                    </Button>
                                </div>

                                {hasValidCode === true && (
                                    <div className="absolute right-0 top-0 mr-28 flex h-full items-center pr-2">
                                        <span className="text-green-500">
                                            <Icons.check className="h-5" />
                                        </span>
                                    </div>
                                )}
                                {hasValidCode === false && code.length > 0 && (
                                    <div className="absolute right-0 top-0 mr-28 flex h-full items-center pr-2">
                                        <span className="text-red-500">
                                            <Icons.close className="h-5" />
                                        </span>
                                    </div>
                                )}
                            </div>
                            {hasValidCode === true && org !== undefined && (
                                <div className="grid gap-2">
                                    <p className='text-xs text-gray-300 dark:text-gray-600'>Welcome, <span className='font-semibold text-gray-500'>{org.name}</span></p>
                                    {isProspective && (
                                        <div className='grid gap-2 gap-y-3 pt-2'>
                                            <Separator />
                                            <Label htmlFor="name">Full Name</Label>
                                            <Input placeholder='John Appleseed' required onChange={(v) => {
                                                setProspectName(v.target.value)
                                            }} />
                                            <Label htmlFor="company">Company</Label>
                                            <Input placeholder='Acme Ventures' required onChange={(v) => {
                                                setProspectCompany(v.target.value)
                                            }} />
                                            <Separator />
                                        </div>
                                    )}
                                </div>
                            )
                            }

                        </div>
                        <div className="grid gap-1">
                            <div className="space-y-1">
                                <Label htmlFor="password">Set Username</Label>
                                <div className="relative">
                                        <Input className=" ring-0" id="deviceLabel" type="text" placeholder="angelo.snr" required minLength={4} maxLength={10} onChange={(e) => {
                                        if (e.target.value == "") {
                                            setIsAvailable(null);
                                        } else {
                                            setAlias(e.target.value);
                                            checkAvailability(e.target.value);
                                        }
                                    }} />
                                    {isAvailable === true && (
                                        <div className="absolute right-0 top-0 flex h-full items-center pr-2">
                                            <span className="text-green-500">
                                                <Icons.check className="h-5" />
                                            </span>
                                        </div>
                                    )}
                                    {isAvailable === false && (
                                        <div className="absolute right-0 top-0 flex h-full items-center pr-2">
                                            <span className="text-red-500">
                                                <Icons.close className="h-5" />
                                            </span>
                                        </div>
                                    )}
                                </div>
                                <div className="py-2">
                                    <Button className="w-full rounded-[2rem]" disabled={checkDisabled()} onClick={(e) => {
                                        onRegister()
                                    }}>
                                        <Icons.fingerprint className="h-6 w-6 pr-1" /> Generate Passkey
                                    </Button>
                                </div>
                            </div>
                        </div>
                        <Separator />
                    </CardContent>
                    <CardFooter className="center justify-center align-middle">
                        <div className="ml-12 grid grid-cols-2 justify-center gap-8 align-middle">
                            <Button variant="ghost" className="w-max font-light" onClick={(e) => {
                                window.open("https://tally.so/r/wdNbdz", "_blank")
                            }}>
                                Need Invite?
                            </Button>
                            <Button variant="link" className="w-max" onClick={(e) => {
                                setIsSignInVisible(true)
                            }}>
                                Sign in
                            </Button>
                        </div>
                    </CardFooter>
                </Card>
            )}
            {isSignInVisible && (
                <LoginView open={isSignInVisible} setOpen={setIsSignInVisible} handleCredentialSet={handleCredentialSet} handleKeygenComplete={handleLoginComplete} />
            )}
        </>
    )
}


interface LoginViewProps {
    open: boolean
    setOpen: (open: boolean) => void
    handleKeygenComplete: (did: string, didDocument: DidDocument) => Promise<void>;
    handleCredentialSet: (credential: PublicKeyCredential) => Promise<void>;
}

function LoginView({ open, setOpen, handleKeygenComplete, handleCredentialSet }: LoginViewProps) {
    const [isAvailable, setIsAvailable] = useState<boolean | null>(null);
    const [alias, setAlias] = useState("");
    const sonr = useSonr()
    const checkAvailability = async (v: string) => {
        const resp = await sonr.checkAlias(v);
        setIsAvailable(!resp.available && v.length >= 4);
    };

    const onLogin = async () => {
        await sonr.login({ alias, onCredentialSet: handleCredentialSet, onLoginComplete: handleKeygenComplete });
    };

    return (
        <Card className="max-w-3xl pb-4 shadow-sm dark:border-slate-600/40 dark:shadow-lg">
            <CardHeader className="space-y-1">
                <CardTitle className="py-2 text-2xl">Login to Sonr</CardTitle>
                <CardDescription>
                    Enter your domain alias to login to Sonr
                </CardDescription>
            </CardHeader>
            <CardContent className="grid gap-4">
                <div className="grid gap-2">
                    <div className="space-y-1">
                        <Label htmlFor="password">Sonr.ID or Address</Label>
                        <div className="relative">
                            <Input autocomplete="username webauthn" className=" ring-0" id="username webauthn" type="username webauthn" placeholder="angelo.snr" required minLength={4} onChange={(e) => {
                                if (e.target.value == "") {
                                    setIsAvailable(null);
                                } else {
                                    setAlias(e.target.value);
                                    checkAvailability(e.target.value);
                                }
                            }} />
                            {isAvailable === true && (
                                <div className="absolute right-0 top-0 flex h-full items-center pr-2">
                                    <span className="text-green-500">
                                        <Icons.check className="h-5" />
                                    </span>
                                </div>
                            )}
                            {isAvailable === false && (
                                <div className="absolute right-0 top-0 flex h-full items-center pr-2">
                                    <span className="text-red-500">
                                        <Icons.close className="h-5" />
                                    </span>
                                </div>
                            )}
                        </div>
                        <div className="py-2">
                            <Button className="w-full rounded-[2rem]" disabled={alias.length < 4 || isAvailable === false} onClick={(e) => {
                                onLogin()
                            }}>
                                <Icons.keyprint className="h-6 w-6 pr-1" /> Continue with Passkey
                            </Button>
                        </div>
                    </div>
                </div>
                <Separator />
            </CardContent>
            <CardFooter>
                <div className="align-end grid grid-cols-2 justify-center gap-8">
                    <Button variant="link" className="w-max" onClick={(e) => {
                        setOpen(false)
                    }}>
                        Return to Sign up
                    </Button>
                </div>
            </CardFooter>
        </Card>
    )
}
