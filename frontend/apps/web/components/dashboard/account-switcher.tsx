"use client"

import * as React from "react"
import { Check, ChevronsUpDown, PlusCircle } from "lucide-react"

import { cn } from "@/lib/utils"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Button } from "@/components/ui/button"
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  CommandSeparator,
} from "@/components/ui/command"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"

import { Sonr } from "@/hooks/useSonr"
import { Account } from "../../../../packages/client/lib/types/user"
import va from "@vercel/analytics"

type PopoverTriggerProps = React.ComponentPropsWithoutRef<typeof PopoverTrigger>

interface TeamSwitcherProps extends PopoverTriggerProps {
  className?: string
  sonr: Sonr
}

export default function AccountSwitcher({ className, sonr }: TeamSwitcherProps) {
  const [open, setOpen] = React.useState(false)
  const [showNewTeamDialog, setShowNewTeamDialog] = React.useState(false)
  const [type, setType] = React.useState('');
  const [name, setName] = React.useState('');

  const [selectedTeam, setSelectedTeam] = React.useState<Account>(
    {
      name: "Account #0",
      did: "did:ethr:0x123456789abcdef123456789abcdef123456789",
      coin_type: "Ethereum",
      chain_id: "1",
      public_key: "0x123456789abcdef123456789abcdef123456789",
      type: "Ethereum",
      address: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
    }
  )
  let blockchainNames = ["Bitcoin", "Ethereum", "Filcoin", "Handshake", "Dogecoin"]
  let accounts = [
    {
      name: "Account #0",
      blockchain: "Bitcoin",
      address: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
      value: "btc",
    },
    {
      name: "Account #1",
      blockchain: "Ethereum",
      address: "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266",
      value: "eth",
    },
    {
      name: "Account #2",
      blockchain: "Filcoin",
      address: "f1z2nm3414214fzsdfgaw4",
      value: "fil",
    },
  ];
  const formatWeb3Address = (address) => `${address.slice(0, 3)}...${address.slice(-8)}`;
  return (
    <Dialog open={showNewTeamDialog} onOpenChange={setShowNewTeamDialog}>
      <Popover open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild>
          <Button
            variant="ghost"
            size="sm"
            role="combobox"
            aria-expanded={open}
            aria-label="Select a account"
            className={cn("w-[200px] justify-between", className)}
          >
            <Avatar className="mr-2 h-5 w-5">
              <AvatarImage
                src={`https://avatar.vercel.sh/${selectedTeam.address}.png`}
                alt={selectedTeam.address}
              />
              <AvatarFallback>SC</AvatarFallback>
            </Avatar>
            <p className="font-mono text-xs">{formatWeb3Address(selectedTeam.address)}</p>
            <ChevronsUpDown className="ml-auto h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </PopoverTrigger>
        <PopoverContent align="end" className="w-fit p-0">
          <Command>
            <CommandList>
              <CommandInput className="m-1 h-min border-0 p-1 text-sm ring-0" placeholder="Search for account..." />
              <CommandEmpty>No account found.</CommandEmpty>
              {blockchainNames.map((group) => (
                <CommandGroup key={group} heading={group}>
                  {accounts.filter((v) => (v.blockchain == group)).map((team) => (
                    <CommandItem
                      key={team.address}
                      value={team.address}
                      onSelect={(v) => {
                        let team = accounts.find((t) => t.address === v)
                        if (!team) return
                        let acc = {
                          name: team.name,
                          did: team.address,
                          coin_type: "",
                          chain_id: "sonr-venus-1",
                          public_key: "",
                          type: "",
                          address: team.address,
                        }
                        setSelectedTeam(acc)
                        setOpen(false)
                        va.track('SelectedAccountChange', { address: acc.address });
                      }}
                      className="text-sm"
                    >
                      <p className="font-mono text-xs">{team.address}</p>
                      <Check
                        className={cn(
                          "ml-2 h-3 w-3",
                          selectedTeam.address === team.address
                            ? "opacity-90"
                            : "opacity-0"
                        )}
                      />
                    </CommandItem>
                  ))}
                </CommandGroup>
              ))}
            </CommandList>
            <CommandSeparator />
            <CommandList>
              <CommandGroup>
                <DialogTrigger asChild>
                  <CommandItem
                    onSelect={() => {
                      setOpen(false)
                      setShowNewTeamDialog(true)
                    }}
                  >
                    <PlusCircle className="mr-2 h-5 w-5" />
                    Create Account
                  </CommandItem>
                </DialogTrigger>
              </CommandGroup>
            </CommandList>
          </Command>
        </PopoverContent>
      </Popover>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Create account</DialogTitle>
          <DialogDescription>
            Add a new crypto account for a blockchain network.
          </DialogDescription>
        </DialogHeader>
        <div>
          <div className="space-y-4 py-2 pb-4">
            <div className="space-y-2">
              <Label htmlFor="account_name">Account name</Label>
              <Input id="name" placeholder="Account #0" onChange={(v) => { setName(v.currentTarget.value) }} />
            </div>
            <div className="space-y-2">
              <Label htmlFor="blockchain">Blockchain</Label>
              <Select onValueChange={(v) => { setType(v) }}>
                <SelectTrigger>
                  <SelectValue placeholder="Select a network" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="bitcoin">
                    <span className="font-medium">Bitcoin</span>
                  </SelectItem>
                  <SelectItem value="ethereum">
                    <span className="font-medium">Ethereum</span>
                  </SelectItem>
                  <SelectItem value="filcoin">
                    <span className="font-medium">Filcoin</span>
                  </SelectItem>
                  <SelectItem value="handshake">
                    <span className="font-medium">Handshake</span>
                  </SelectItem>
                  <SelectItem value="dogecoin">
                    <span className="font-medium">Dogecoin</span>
                  </SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>
        </div>
        <DialogFooter className="flex gap-x-2">
          <Button variant="outline" onClick={() => setShowNewTeamDialog(false)}>
            Cancel
          </Button>
          <Button type="submit" onClick={async () => {
            const account = await sonr.client.accounts.create(name, type)
            console.log(account)
          }}>Confirm</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
