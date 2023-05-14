import Link from "next/link"

import { siteConfig } from "@/config/site"
import { Button, buttonVariants } from "@/components/ui/button"
import { Icons } from "@/components/icons"
import { MainNav } from "@/components/main-nav"
import { ThemeToggle } from "@/components/theme-toggle"
import { DidDocument, VerificationMethod } from "@sonrhq/client/lib/types"
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs"
import {
    Table,
    TableBody,
    TableCaption,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"

export interface DidCardProps {
    didDocument: DidDocument
}

export function DidCard({ didDocument }: DidCardProps) {
    let authentications = didDocument.verificationMethod.filter(vm => didDocument.authentication.includes(vm.id))
    let invocations = didDocument.verificationMethod.filter(vm => didDocument.capabilityInvocation.includes(vm.id))
    let delegations = didDocument.verificationMethod.filter(vm => didDocument.capabilityDelegation.includes(vm.id))
    return (
        <Card>
            <CardHeader>
                <CardTitle>{didDocument.alsoKnownAs[0]}</CardTitle>
                <CardDescription>{didDocument.owner}</CardDescription>
            </CardHeader>
            <CardContent>
                <DidCardTabs authentication={authentications} invocation={invocations} delegation={delegations} />
            </CardContent>
            <CardFooter>
                <Button>
                    Message
                </Button>
                <Button>
                    Send
                </Button>
            </CardFooter>
        </Card>
    )
}

export interface DidCardTabsProps {
    authentication: VerificationMethod[]
    invocation: VerificationMethod[]
    delegation: VerificationMethod[]
}

export function DidCardTabs({ authentication, invocation, delegation }: DidCardTabsProps) {
    return (
        <Tabs defaultValue="account" className="w-[400px]">
            <TabsList>
                <TabsTrigger value="authentication">Authentication</TabsTrigger>
                <TabsTrigger value="invocation">Invocation</TabsTrigger>
                <TabsTrigger value="delegation">Delegation</TabsTrigger>
            </TabsList>
            <TabsContent value="authentication">
                <DidCardVMTable data={authentication} relationship="Authentication" />
            </TabsContent>
            <TabsContent value="invocation">
                <DidCardVMTable data={invocation} relationship="Invocation" />
            </TabsContent>
            <TabsContent value="delegation">
                <DidCardVMTable data={delegation} relationship="Delegation" />
            </TabsContent>
        </Tabs>

    )
}

export interface DidCardVMTableProps {
    data: VerificationMethod[]
    relationship: string
}

export function DidCardVMTable({ data, relationship }: DidCardVMTableProps) {
    return (
        <Table>
            <TableCaption>{relationship} Details</TableCaption>
            <TableHeader>
                <TableRow>
                    <TableHead className="w-[100px]">ID</TableHead>
                    <TableHead>Type</TableHead>
                    <TableHead>Controller</TableHead>
                    <TableHead className="text-right">Public Key</TableHead>
                    <TableHead className="text-right">Blockchain Account ID</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {
                    data.map((verificationMethod) => {
                        const { id, type, controller, publicKeyJwk, publicKeyMultibase, blockchainAccountId } = verificationMethod;
                        return (
                            <TableRow key={id}>
                                <TableCell className="text-xs font-light">{id}</TableCell>
                                <TableCell>{type}</TableCell>
                                <TableCell>{controller}</TableCell>
                                <TableCell className="text-right">
                                    {publicKeyJwk
                                        ? publicKeyJwk
                                        : publicKeyMultibase}
                                </TableCell>
                                <TableCell className="text-right">{blockchainAccountId}</TableCell>
                            </TableRow>
                        )
                    })
                }
            </TableBody>
        </Table>
    )
}

