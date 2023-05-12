
import React from "react"
import { CommandDialog, CommandInput, CommandList, CommandEmpty, CommandGroup, CommandItem } from "@/components/ui/command"
import { cmdConfig } from "@/config/cmd"
import Image from "next/image";

interface CommandMenuProps {
    open: boolean
    setOpen: React.Dispatch<React.SetStateAction<boolean>>
}

const CommandMenu = ({ open, setOpen }: CommandMenuProps) => {
    const [filter, setFilter] = React.useState("")

    const filteredCmds = React.useMemo(() => {
        const filteredGroups = Object.entries(cmdConfig.cmds).map(([groupName, groupItems]) => {
            const filteredItems = groupItems.filter(item => item.name.toLowerCase().includes(filter.toLowerCase()))
            if (filteredItems.length === 0) {
                return null
            }
            return (
                <CommandGroup key={groupName} heading={groupName}>
                    {filteredItems.map(item => (
                        <CommandItem key={item.name} {...item} />
                    ))}
                </CommandGroup>
            )
        })
        return filteredGroups.filter(Boolean)
    }, [filter])

    React.useEffect(() => {
        const down = (e: KeyboardEvent) => {
            if (e.key === "k" && e.metaKey) {
                setOpen(open => !open)
            }
        }
        document.addEventListener("keydown", down)
        return () => document.removeEventListener("keydown", down)
    }, [setOpen])

    return (
        <CommandDialog open={open} onOpenChange={setOpen}>
            <CommandInput placeholder={cmdConfig.help} onValueChange={e => setFilter(e)} />
            <CommandList className="w-[340px]">
                {Object.entries(cmdConfig.cmds).map(([groupName, groupItems]) => {
                    return (
                        <CommandGroup key={groupName}>
                            <p className="text-xs font-light text-gray-500">{groupName}</p>
                            {groupItems.map(item => (
                                <CommandItem key={item.name}> <Image className="mr-2" width={20} height={20} src={item.icon} alt="Command Icon" /> {item.name} </CommandItem>
                            ))}
                </CommandGroup>
                    )
                })}

            </CommandList>
        </CommandDialog>
    )
}

export { CommandMenu }
