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
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"
import { Textarea } from "@/components/ui/textarea"

export function DemoReportAnIssue() {
    return (
        <Card className="col-span-4">
            <CardHeader>
                <CardTitle>Send Payment</CardTitle>
                <CardDescription>
                    What area are you having problems with?
                </CardDescription>
            </CardHeader>
            <CardContent className="grid gap-6">
                <div className="grid gap-2">
                    <Label htmlFor="subject">Amount</Label>
                    <Input id="subject" type="number" placeholder="$400.00" />
                </div>
                <div className="grid gap-2">
                    <Label htmlFor="description">Memo</Label>
                    <Textarea
                        id="description"
                        placeholder="The Times 03/Jan/2009 Chancellor on brink of second bailout for banks."
                    />
                </div>
            </CardContent>
            <CardFooter className="justify-between space-x-2">
                <Button variant="ghost">Cancel</Button>
                <Button>Submit</Button>
            </CardFooter>
        </Card>
    )
}
