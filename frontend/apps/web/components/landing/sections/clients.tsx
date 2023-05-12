import { AlanLogo } from "@/components/landing/logo/alan";
import { ArcLogo } from "@/components/landing/logo/arc";
import { CashAppLogo } from "@/components/landing/logo/cashapp";
import { DescriptLogo } from "@/components/landing/logo/descript";
import { LoomLogo } from "@/components/landing/logo/loom";
import { MercuryLogo } from "@/components/landing/logo/mercury";
import { OpenSeaLogo } from "@/components/landing/logo/opensea";
import { PitchLogo } from "@/components/landing/logo/pitch";
import { RampLogo } from "@/components/landing/logo/ramp";
import { RaycastLogo } from "@/components/landing/logo/raycast";
import { RetoolLogo } from "@/components/landing/logo/retool";
import { VercelLogo } from "@/components/landing/logo/vercel";


export const Clients = () => (
    <>
        <p className="mb-12 text-center text-lg text-white md:text-xl">
            <span className="text-primary-text">
                Powering the world’s best product teams.
            </span>
            <br className="hidden md:block" /> From next-gen startups to established
            enterprises.
        </p>

        <div className="flex flex-wrap justify-around gap-x-6 gap-y-8 [&_svg]:max-w-[16rem] [&_svg]:basis-[calc(50%-12px)] md:[&_svg]:basis-[calc(16.66%-20px)]">
            {/* <RampLogo />
            <LoomLogo className="hidden md:block" />
            <VercelLogo />
            <DescriptLogo className="hidden md:block" />
            <CashAppLogo />
            <RaycastLogo />
            <MercuryLogo />
            <RetoolLogo />
            <AlanLogo className="hidden md:block" />
            <ArcLogo className="hidden md:block" />
            <OpenSeaLogo className="hidden md:block" />
            <PitchLogo className="hidden md:block" /> */}
        </div>
    </>
);
