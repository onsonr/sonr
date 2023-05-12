import React from "react"
import { DialogTitle, DialogDescription, DialogHeader } from "@/components/ui/dialog"



const RegisterDialog = ({ open, setOpen }: { open: boolean, setOpen: (open: boolean) => void }) => {
    const [formDone, setFormDone] = React.useState(false)
    const [formValues, setFormValues] = React.useState({
        deviceLabel: "",
        pin: "",
        pinConfirm: ""
    })

    return (<DialogHeader className="mb-0">
        <DialogTitle> <h2 className="mt-10 scroll-m-20 border-b border-b-slate-200 pb-2 text-2xl font-semibold tracking-tight transition-colors first:mt-0 dark:border-b-slate-700">
            Register New ID
        </h2></DialogTitle>
        <DialogDescription>
            Your privacy is paramount, and our state-of-the-art authentication ensures that your data remains protected.
        </DialogDescription>
    </DialogHeader>
    )
}
