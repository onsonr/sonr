import Image from 'next/image'

import { Button } from '@/components/Button'
import { Heading } from '@/components/Heading'


const libraries = [
  {
    href: '#',
    name: 'Go',
    description:
      'Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.',
    logo: "/images/go-original.svg",
  },
  {
    href: '#',
    name: 'React',
    description:
      'A JavaScript library for building user interfaces. React is a declarative, efficient, and flexible JavaScript library for building user interfaces.',
    logo: "/images/react-original.svg",
  },
  {
    href: '#',
    name: 'Flutter',
    description:
      'Flutter allows building beautiful, natively compiled applications for mobile, web, and desktop from a single codebase.',
    logo: "/images/flutter-original.svg",
  },
]

const communityLibraries = [
  {
    href: '#',
    name: 'Swift',
    description:
      'Swift is a general-purpose programming language built using a modern approach to safety, performance, and software design patterns.',
    logo: "/images/swift-original.svg",
  },
  {
    href: '#',
    name: 'Java',
    description:
      'Java implementation for Android and other Java-based platforms.',
    logo: "/images/java-original.svg",
  }
]


export function Libraries() {
  return (
    <div className="my-16 xl:max-w-none">
      <Heading level={2} id="official-libraries">
        Official libraries
      </Heading>
      <div className="not-prose mt-4 grid grid-cols-1 gap-x-6 gap-y-10 border-t border-zinc-900/5 pt-10 dark:border-white/5 sm:grid-cols-2 xl:max-w-none xl:grid-cols-3">
        {libraries.map((library) => (
          <div key={library.name} className="flex flex-row-reverse gap-6">
            <div className="flex-auto">
              <h3 className="text-sm font-semibold text-zinc-900 dark:text-white">
                {library.name}
              </h3>
              <p className="mt-1 text-sm text-zinc-600 dark:text-zinc-400">
                {library.description}
              </p>
              <p className="mt-4">
                <Button href={library.href} variant="text" arrow="right">
                  Read more
                </Button>
              </p>
            </div>
            <Image
              height={10}
              width={10}
              src={library.logo}
              alt=""
              className="h-12 w-12"
              unoptimized
            />
          </div>
        ))}
      </div>
    </div>
  )
}

export function CommunityLibraries() {
  return (
    <div className="my-16 xl:max-w-none">
      <Heading level={2} id="official-libraries">
        Community libraries
      </Heading>
      <div className="not-prose mt-4 grid grid-cols-1 gap-x-6 gap-y-10 border-t border-zinc-900/5 pt-10 dark:border-white/5 sm:grid-cols-2 xl:max-w-none xl:grid-cols-3">
        {communityLibraries.map((library) => (
          <div key={library.name} className="flex flex-row-reverse gap-6">
            <div className="flex-auto">
              <h3 className="text-sm font-semibold text-zinc-900 dark:text-white">
                {library.name}
              </h3>
              <p className="mt-1 text-sm text-zinc-600 dark:text-zinc-400">
                {library.description}
              </p>
              <p className="mt-4">
                <Button href={library.href} variant="text" arrow="right">
                  Read more
                </Button>
              </p>
            </div>
            <Image
              height={10}
              width={10}
              src={library.logo}
              alt=""
              className="h-12 w-12"
              unoptimized
            />
          </div>
        ))}
      </div>
    </div>
  )
}
