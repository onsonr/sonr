"use client"
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
} from "@/components/ui/dialog"
import { Label } from "@/components/ui/label"
import React from "react"
import QRCode from "react-qr-code"


interface QRDialogProps {
    open: boolean
    setOpen: (open: boolean) => void
    address: string
}

const QRDialog = ({ open, setOpen, address }: QRDialogProps) => (
    <Dialog onOpenChange={setOpen} open={open}>
        <DialogContent className="flex-auto p-8">
            <DialogHeader className="mb-0">
                <DialogTitle>
                    <h1 className="mt-10 scroll-m-20 border-b border-b-slate-200 pb-2 text-3xl font-extrabold tracking-tight transition-colors first:mt-0 dark:border-b-slate-700">
                        Send or Receive
                    </h1>
                </DialogTitle>

                <DialogDescription>
                    <div style={{ height: "auto", margin: "0 auto", maxWidth: 256, width: "100%" }} className="py-4">
                        <QRCode
                            size={256}
                            style={{ height: "auto", maxWidth: "100%", width: "100%" }}
                            value={address}
                            viewBox={`0 0 256 256`}
                        />
                    </div>
                    <div className="p-2" />
                    <Label className="ml-4 mt-4 rounded-md bg-slate-900 px-4 py-2 text-center text-white">{address}</Label>
                </DialogDescription>
            </DialogHeader>
        </DialogContent>
    </Dialog >
)

export { QRDialog }
