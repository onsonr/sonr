"use client"

import Link from "next/link"

import { cn } from "@/lib/utils"

export function DashboardNav({
  className,
  ...props
}: React.HTMLAttributes<HTMLElement>) {
  return (
    <nav
      className={cn("flex items-center space-x-6", className)}
      {...props}
    >
      <Link
        href="/wallet"
        className="text-muted-foreground hover:text-primary text-sm font-medium transition-colors"
      >
        Wallet
      </Link>
      <Link
        href="/profile"
        className="hover:text-primary text-sm font-medium transition-colors"
      >
        Explorer
      </Link>
      <Link
        href="/roadmap"
        className="text-muted-foreground hover:text-primary text-sm font-medium transition-colors"
      >
        Roadmap
      </Link>
    </nav>
  )
}
