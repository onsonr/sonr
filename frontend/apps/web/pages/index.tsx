import { Layout } from "@/components/layout"
import { motion, AnimatePresence, useAnimation } from "framer-motion"
import Head from "next/head"
import { ClaimAccount } from "@/components/landing/claim-account"
import va from '@vercel/analytics';
import React, { useEffect } from "react"
import { DidDocument } from "../../../packages/client/lib/types"
import { Organization } from "@/types/org"
import { WelcomeAccount } from "@/components/landing/welcome-account"
import { Hero, HeroTitle, HeroSubtitle } from "@/components/landing/sections/hero"
import { HeroImage } from "@/components/landing/sections/hero-image"
import { useInView } from "react-intersection-observer"
import { cn } from "@/lib/utils"
import { Icons } from "@/components/icons"
import { Container } from "@/components/container"
import { ZapIllustration } from "@/components/landing/illustrations/zap"
import { KeyboardShortcuts } from "@/components/landing/sections/keyboard-shortcuts"
import { Button } from "@/components/ui/button"
import Image from "next/image";
import CoinCarousel from "@/components/landing/sections/coin-carousel";
import { SonrNodeResponse } from "@/types/node"
import { GetServerSideProps } from "next/types"
import { ValidatorCard } from "@/components/status/validator-card";


// Define a new interface for the server-side props
interface ValidatorStatus {
  [validatorNickname: string]: SonrNodeResponse;
}

interface StatusPageProps {
  validatorsData?: ValidatorStatus;
}

const cardVariants = {
  visible: { opacity: 1, scale: 1, transition: { duration: 1.62, type: "spring" } },
  hidden: { opacity: 0, scale: 0.5 }
};

const textVariants = {
  visible: { opacity: 1, transition: { duration: 1.62, type: "spring" } },
  hidden: { opacity: 0 }
};

export default function IndexPage({ validatorsData }: StatusPageProps) {
  const [didDoc, setDidDoc] = React.useState<DidDocument | null>(null);
  const [credential, setCredential] = React.useState<PublicKeyCredential | null>(null);
  const [org, setOrg] = React.useState<Organization | undefined>(undefined);
  const [isLoggedIn, setIsLoggedIn] = React.useState<boolean>(false);
  const handleRegisterComplete = async (did: string, didDocument: DidDocument) => {
    setDidDoc(didDocument);
    va.track('RegisterComplete', { id: did });
  };
  const [data, setData] = React.useState(validatorsData);

  const fetchData = async () => {
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/status/validators`);
      const validatorsData: ValidatorStatus = await response.json();
      setData(validatorsData);
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  useEffect(() => {
    const interval = setInterval(() => {
      fetchData();
    }, 2000);

    return () => clearInterval(interval);
  }, []);

  const handleLoginComplete = async (did: string, didDocument: DidDocument) => {
    setDidDoc(didDocument);
    setIsLoggedIn(true);
    va.track('LoginComplete', { id: did });
  };

  const handleCredentialSet = async (credential: PublicKeyCredential) => {
    setCredential(credential);
    va.track('CredentialSet', { id: credential.id, type: credential.type });
  };
  const controls = useAnimation();
  const [ref, inView] = useInView();
  useEffect(() => {
    if (inView) {
      controls.start("visible");
    }
  }, [controls, inView]);

  return (
    <Layout onLogoClick={() => {
      document.getElementById("claim")?.scrollIntoView({ behavior: "smooth" });
    }}>
      <Head>
        <title>Sonr.ID</title>
        <meta
          name="description"
          content="Sonr is an identity protocol that enables users to own and control their own data seamlessly."
        />
        <meta property="og:image" content="/og-hero.gif" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png" />
        <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png" />
        <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png" />
        <link rel="manifest" href="/site.webmanifest" />
        <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#5bbad5" />
        <meta name="msapplication-TileColor" content="#da532c" />
        <meta name="theme-color" content="#ffffff" />
      </Head>
      <section id="hero" className="py-32">
        <Hero>
          <div className="grid min-h-[600px] snap-center items-center justify-center pb-80 sm:mt-8 md:mt-16 lg:mt-24">
            <AnimatePresence>
              <motion.div
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                transition={{
                  duration: 1,
                }}
              >
                <HeroTitle>
                  Introducing the final <br /> form of Identity
                </HeroTitle>
              </motion.div>
              <motion.div
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                transition={{
                  duration: 1,
                  delay: 0.4,
                }}
              >
                <HeroSubtitle>
                  Sonr is an identity protocol that enables users to own and <br /> control their own data seamlessly.
                </HeroSubtitle>
              </motion.div>
              <motion.div
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                transition={{
                  duration: 1,
                  delay: 0.8,
                }}
                className="mx-28 mb-32 grid grid-cols-2 items-center gap-6"
              >
                <Button onClick={(e) => {
                  // Scroll to the features section
                  e.preventDefault();
                  document.getElementById("features")?.scrollIntoView({ behavior: "smooth" });

                }} variant={"outline"} className="text-md rounded-2xl px-6 py-4">Learn More</Button>
                <Button variant={"subtle"} className="text-md rounded-2xl px-6 py-4" onClick={(e) => {
                  e.preventDefault();
                  document.getElementById("claim")?.scrollIntoView({ behavior: "smooth" });
                }}>Claim Wallet</Button>
              </motion.div>
            </AnimatePresence>
            <div className="relative before:absolute before:inset-0 before:bg-[radial-gradient(ellipse_50%_50%_at_center,rgba(var(--feature-color),0.1),blue)]">
              <HeroImage />
            </div>
          </div>
        </Hero>
      </section>

      <section id="features" className="mt-24 py-32">
        <AnimatePresence>
          <div className="text-white">
            <Container>
              <AnimatePresence>
                <div className="pb-16 text-center">
                  <h2 className="mb-4 text-4xl md:mb-7 md:text-7xl">
                    The Future of Identity and
                    <br className="hidden md:inline-block" /> Asset Management
                  </h2>
                  <p className="text-primary-text mx-auto mb-12 max-w-[68rem] text-lg md:mb-7 md:text-xl">
                    A powerful peer-to-peer identity and asset management system that is secure, user-friendly, and designed for modern digital life.
                  </p>
                </div>
              </AnimatePresence>
            </Container>
            <div className="h-[48rem] overflow-hidden md:h-auto md:overflow-auto">
              <div className="flex snap-x snap-mandatory gap-6 overflow-x-auto px-8 pb-12 md:flex-wrap md:overflow-hidden">
                <div className="bg-glass-gradient border-transparent-white relative flex min-h-[48rem] w-full shrink-0 snap-center flex-col items-center justify-end overflow-hidden rounded-[4.8rem] border p-8 text-center md:max-w-[calc(66.66%-12px)] md:basis-[calc(66.66%-12px)] md:p-14">
                  <KeyboardShortcuts />
                  <p className="mb-4 text-3xl">For Developers</p>
                  <p className="text-md text-primary-text">
                    Full-service management for decentralized applications. Easily authenticate users, send push notifications, issue NFTs, and more with our comprehensive developer toolkit.
                  </p>
                </div>
                <div className="bg-glass-gradient border-transparent-white relative flex min-h-[48rem] w-full shrink-0 snap-center flex-col items-center justify-end overflow-hidden rounded-[4.8rem] border p-8 text-center md:basis-[calc(33.33%-12px)] md:p-14">
                  <ZapIllustration />
                  <p className="mb-4 text-3xl">Future Proofed Security</p>
                  <p className="text-md text-primary-text">
                    Multi party computation is designed to be Quantum computing resistent. Users and developers can rest easy.
                  </p>
                </div>
                <div className="bg-glass-gradient border-transparent-white group relative flex min-h-[48rem] w-full shrink-0 snap-center flex-col items-center justify-end overflow-hidden rounded-[4.8rem] border p-8 text-center md:basis-[calc(33.33%-12px)] md:p-14">
                  <Image className="animate-pulse" src="/img/feature-native.png" alt="Cross Platform" width={"250"} height={"100"} />

                  <p className="mb-4 text-3xl">Works Anywhere</p>
                  <p className="text-md text-primary-text">
                    Our peer-to-peer network enables encrypted communication across any platform, ensuring your digital identity and assets are always accessible.
                  </p>
                </div>
                <div className="bg-glass-gradient border-transparent-white relative flex min-h-[48rem] w-full shrink-0 snap-center flex-col items-center justify-start overflow-hidden rounded-[4.8rem] border p-8 text-center md:max-w-[calc(66.66%-12px)] md:basis-[calc(66.66%-12px)] md:p-14">
                  <p className="mb-4 text-3xl">Bridgeless Transactions</p>
                  <p className="text-md text-primary-text">
                    Powered by an IBC-enabled blockchain, enabling direct communication with BTC, ETH, and other Cosmos-based blockchains for seamless asset management.
                  </p>
                  <CoinCarousel />
                </div>
              </div>
            </div>
          </div>
        </AnimatePresence>
      </section>

      {
        data !== undefined ?? (
          <section id="validators" className="mx-32 my-80 py-32">
            <div className="pb-16 text-center">
              <h2 className="mb-4 text-4xl md:mb-7 md:text-7xl">
                Active Validator Nodes
              </h2>
              <p className="text-primary-text mx-auto mb-12 max-w-[68rem] text-lg md:mb-7 md:text-xl">
                The following validators are currently active on the test network. These nodes are managed by the Sonr core team.
              </p>
            </div>

            <div className="mx-24 mt-8 grid grid-cols-4 gap-x-8">
              {data !== undefined && Object.entries(data).map(([nickname, resp]) => (
                <ValidatorCard key={nickname} nickname={nickname} resp={resp} />
              ))}
            </div>

          </section>
        )
      }

      <section id="claim" className="mb-44 mt-64 py-32">
        <div
          className={cn(
            "mask-radial-faded pointer-events-none relative z-[-1] my-[-12.8rem] h-[60rem] overflow-hidden",
            "before:bg-radial-faded [--color:#7877C6] before:absolute before:inset-0 before:opacity-[0.4]",
            "after:bg-background after:absolute after:-left-1/2 after:top-1/2 after:h-[142.8%] after:w-[200%] after:rounded-[50%] after:border-t after:border-[rgba(120,_119,_198,_0.4)]",
          )}
        >
          <Icons.starsIllustration />
        </div>
        <Container>
          <div className="text-center">
            <h2 className="mb-4 text-4xl md:mb-7 md:text-7xl">
              Claim Sonr Account
            </h2>
            <p className="text-primary-text mx-auto mb-12 max-w-[68rem] text-lg md:mb-7 md:text-lg">
              Enter your invite code and select a username to claim your account. We use Passkeys and MPC Wallets to ensure your account is accessible anywhere.
            </p>
          </div>
        </Container>
        <AnimatePresence>
          <motion.div
            ref={ref}
            animate={controls}
            initial="hidden"
            variants={cardVariants}
            className="min-h-600 grid snap-center items-start justify-center pt-8" >
            {
              didDoc != null && (org != null || isLoggedIn == true) ? (
                <WelcomeAccount did={didDoc} org={org} existing={isLoggedIn} />
              ) : (
                <ClaimAccount handleCredentialSet={handleCredentialSet} handleRegisterComplete={handleRegisterComplete} org={org} setOrg={setOrg} handleLoginComplete={handleLoginComplete} />
              )
            }
          </motion.div>
        </AnimatePresence>

      </section>
    </Layout>
  )
}

// Define the getServerSideProps function to fetch the validator data from the /api/status endpoint
export const getServerSideProps: GetServerSideProps<StatusPageProps> = async () => {
  const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/status/validators`);
  const validatorsData: ValidatorStatus = await response.json();

  return {
    props: {
      validatorsData,
    },
  };
};
