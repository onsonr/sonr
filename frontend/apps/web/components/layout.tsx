import { SiteHeader } from "@/components/site-header"
import React, { useEffect } from "react"
import { CommandMenu } from "@/components/dashboard/command-menu"
import { Footer } from "@/components/footer"

interface LayoutProps {
  children: React.ReactNode
  onLogoClick?: () => void
}



export function Layout({ children, onLogoClick }: LayoutProps) {
  return (
    <div className="zoom">
      <SiteHeader onLogoClick={onLogoClick} />
      <main className="small-view overflow-hidden align-middle xl:px-[12%]">{children}</main>
      <Footer />
    </div>
  )
}
