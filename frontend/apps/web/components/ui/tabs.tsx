"use client"

import * as React from "react"
import * as TabsPrimitive from "@radix-ui/react-tabs"

import { cn } from "@/lib/utils"

const Tabs = TabsPrimitive.Root

const TabsList = React.forwardRef<
  React.ElementRef<typeof TabsPrimitive.List>,
  React.ComponentPropsWithoutRef<typeof TabsPrimitive.List>
>(({ className, ...props }, ref) => (
  <TabsPrimitive.List
    ref={ref}
    className={cn(
      "inline-flex items-center justify-center rounded-3xl bg-slate-100 px-0.5 py-1.5 dark:border dark:border-gray-700/50 dark:bg-black/30 dark:drop-shadow-lg",
      className
    )}
    {...props}
  />
))
TabsList.displayName = TabsPrimitive.List.displayName

const TabsTrigger = React.forwardRef<
  React.ElementRef<typeof TabsPrimitive.Trigger>,
  React.ComponentPropsWithoutRef<typeof TabsPrimitive.Trigger>
>(({ className, ...props }, ref) => (
  <TabsPrimitive.Trigger
    className={cn(
      "ring-sonr/50 data-[state=active]:shadow-xs mx-1 inline-flex min-w-[100px] items-center justify-center rounded-2xl px-3 py-1.5 text-sm font-medium text-slate-700  transition-all disabled:pointer-events-none disabled:opacity-50 data-[state=active]:rounded-3xl data-[state=active]:bg-white data-[state=active]:text-slate-900 dark:text-slate-200 dark:data-[state=active]:bg-slate-800/80 dark:data-[state=active]:text-slate-100 dark:data-[state=active]:ring-2",
      className
    )}
    {...props}
    ref={ref}
  />
))
TabsTrigger.displayName = TabsPrimitive.Trigger.displayName

const TabsContent = React.forwardRef<
  React.ElementRef<typeof TabsPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof TabsPrimitive.Content>
>(({ className, ...props }, ref) => (
  <TabsPrimitive.Content
    className={cn(
      "mt-2 rounded-md",
      className
    )}
    {...props}
    ref={ref}
  />
))
TabsContent.displayName = TabsPrimitive.Content.displayName

export { Tabs, TabsList, TabsTrigger, TabsContent }
