"use client"

import React, { useEffect, useState } from "react"
import { Logo } from "@/components/logo"

interface SiteHeaderProps {
  onLogoClick?: () => void
}

export function SiteHeader({ onLogoClick }: SiteHeaderProps) {
  return (
    <header className="sticky top-0 z-40 w-full border-b border-b-slate-200 bg-white dark:border-b-slate-900 dark:bg-black">
      <div className="container flex h-16 items-center space-x-4 sm:justify-between sm:space-x-0">

        <nav
          className={"flex items-center space-x-4 lg:space-x-6"}
        >
          <div onClick={(e) => {
            e.preventDefault()
            onLogoClick?.()
          }}>
            <Logo />
          </div>
        </nav>
      </div>
    </header>
  )
}
