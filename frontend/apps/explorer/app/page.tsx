import Link from "next/link"

import { siteConfig } from "@/config/site"
import { buttonVariants } from "@/components/ui/button"
import { ComboboxDemo } from "@/components/search-box"

export default function IndexPage() {
  return (
    <section className="container grid items-center gap-8 py-16 md:py-40 lg:py-64">
      <h1 className="text-3xl font-extrabold leading-tight tracking-tighter sm:text-3xl md:text-5xl lg:text-6xl">
        Sonr Explorer
      </h1>
      <div className="flex gap-4">
        <ComboboxDemo />
      </div>
    </section>
  )
}
