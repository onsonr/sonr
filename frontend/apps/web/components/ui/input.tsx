import * as React from "react"

import { cn } from "@/lib/utils"

export interface InputProps
  extends React.InputHTMLAttributes<HTMLInputElement> {
  autocomplete?: string
}

const Input = React.forwardRef<HTMLInputElement, InputProps>(
  ({ className, autoComplete, ...props }, ref) => {
    return (
      <input
        autoComplete={autoComplete}
        className={cn(
          "rounded-2xl flex h-10 w-full border border-slate-300 bg-transparent px-3 py-2 text-sm placeholder:text-slate-400 focus:outline-none focus:ring-2 focus:ring-slate-400 focus:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 dark:border-gray-700/70 dark:text-slate-50 dark:focus:ring-slate-800 dark:focus:ring-offset-slate-900",
          className
        )}
        ref={ref}
        {...props}
      />
    )
  }
)
Input.displayName = "Input"

export { Input }
