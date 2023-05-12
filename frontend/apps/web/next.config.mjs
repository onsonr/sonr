/** @type {import('next').NextConfig} */

const nextConfig = {
  reactStrictMode: true,
  transpilePackages: ['@sonrhq/ts-client'],
  // Adds unsplash to next.config host list
  images: {
    domains: ["images.unsplash.com"],
  },
  experimental: {
    fontLoaders: [
      {
        loader: "@next/font/google",
        options: { subsets: ["latin"] },
      },
    ],
  },
}

export default nextConfig
