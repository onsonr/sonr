"use client"

import * as React from "react"
import { Check, ChevronsUpDown, Search } from "lucide-react"

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
    Command,
    CommandEmpty,
    CommandGroup,
    CommandInput,
    CommandItem,
} from "@/components/ui/command"
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from "@/components/ui/popover"

import { Input } from "@/components/ui/input"

// Mock DIDs for demo
const dids = [
    {
        value: "did:example:123",
        label: "did:example:123",
    },
    {
        value: "did:example:456",
        label: "did:example:456",
    },
    // More DIDs...
]

export function DidSearchBox() {
    const [open, setOpen] = React.useState(false)
    const [value, setValue] = React.useState("")
    const [input, setInput] = React.useState("")

    const searchResults = dids.filter(did => did.label.includes(input))

    return (
        <Popover open={open} onOpenChange={setOpen}>
            <PopoverTrigger asChild>
                <div className="flex w-full max-w-sm items-center space-x-2">
                    <Input
                        type="text"
                        placeholder="Search DIDs..."
                        value={input}
                        onChange={(e) => setInput(e.target.value)}
                        onFocus={() => setOpen(true)}
                    />
                    <Button type="submit" onClick={() => setOpen(!open)}>
                        <Search className="h-4 w-4" />
                    </Button>
                </div>
            </PopoverTrigger>
            <PopoverContent className="w-[200px] p-0">
                <Command>
                    <CommandInput
                        placeholder="Search DIDs..."
                        value={input}
                        onValueChange={(e) => setInput(e)}
                    />
                    <CommandEmpty>No DID found.</CommandEmpty>
                    <CommandGroup>
                        {searchResults.map((did) => (
                            <CommandItem
                                key={did.value}
                                onSelect={(currentValue) => {
                                    setValue(currentValue === value ? "" : currentValue)
                                    setOpen(false)
                                }}
                            >
                                <Check
                                    className={cn(
                                        "mr-2 h-4 w-4",
                                        value === did.value ? "opacity-100" : "opacity-0"
                                    )}
                                />
                                {did.label}
                            </CommandItem>
                        ))}
                    </CommandGroup>
                </Command>
            </PopoverContent>
        </Popover>
    )
}
