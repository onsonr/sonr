"use client";

import { CryptoCard, CryptoCardLarge } from "@/components/card-crypto";
import { Icons } from "@/components/icons";
import { AspectRatio } from "@radix-ui/react-aspect-ratio";
import classNames from "classnames";
import { CSSProperties, useEffect, useRef, useState } from "react";
import { useInView } from "react-intersection-observer";

const randomNumberBetween = (min: number, max: number) => {
    return Math.floor(Math.random() * (max - min + 1) + min);
};

interface Line {
    id: string;
    direction: "to top" | "to left";
    size: number;
    duration: number;
}

export const HeroImage = () => {
    const { ref, inView } = useInView({ threshold: 0.4, triggerOnce: true });
    const [lines, setLines] = useState<Line[]>([]);
    const timeoutRef = useRef<ReturnType<typeof setTimeout> | null>(null);

    const removeLine = (id: string) => {
        setLines((prev) => prev.filter((line) => line.id !== id));
    };

    useEffect(() => {
        if (!inView) return;

        const renderLine = (timeout: number) => {
            timeoutRef.current = setTimeout(() => {
                setLines((lines) => [
                    ...lines,
                    {
                        direction: Math.random() > 0.5 ? "to top" : "to left",
                        duration: randomNumberBetween(1300, 3500),
                        size: randomNumberBetween(10, 30),
                        id: Math.random().toString(36).substring(7),
                    },
                ]);

                renderLine(randomNumberBetween(700, 2500));
            }, timeout);
        };

        renderLine(randomNumberBetween(700, 1300));

        return () => {
            if (timeoutRef.current) clearTimeout(timeoutRef.current);
        };
    }, [inView, setLines]);

    return (
        <AspectRatio>


            <div ref={ref} className="mt-[12.8rem] [perspective:700px]">
            <div
                className={classNames(
                    "relative rounded-3xl border border-transparent-white bg-white bg-opacity-[0.1] bg-hero-gradient dark:bg-slate-800 dark:bg-opacity-40 w-[700px]",
                    inView ? "animate-image-rotate" : "[transform:rotateX(25deg)]",
                    "before:absolute before:top-0 before:left-0 before:h-full before:w-full before:bg-hero-glow before:opacity-0 before:[filter:blur(120px)]",
                    inView && "before:animate-image-glow"
                )}
            >
                <div className="absolute right-0 top-0 z-20 h-full w-full">
                    {lines.map((line) => (
                        <span
                            key={line.id}
                            onAnimationEnd={() => removeLine(line.id)}
                            style={
                                {
                                    "--direction": line.direction,
                                    "--size": line.size,
                                    "--animation-duration": `${line.duration}ms`,
                                } as CSSProperties
                            }
                            className={classNames(
                                "absolute top-0 block h-[0.25px] w-[10rem] bg-sonr",
                                line.direction === "to left" &&
                                `left-0 h-[1px] w-[calc(var(--size)*0.5rem)] animate-glow-line-horizontal md:w-[calc(var(--size)*1rem)]`,
                                line.direction === "to top" &&
                                `right-0 h-[calc(var(--size)*0.5rem)] w-[1px] animate-glow-line-vertical md:h-[calc(var(--size)*1rem)]`
                            )}
                        />
                    ))}
                </div>
                <svg
                    className={classNames(
                        "absolute left-0 top-0 h-full w-full",
                        "[&_path]:stroke-blue-100 [&_path]:[strokeOpacity:0.4] [&_path]:[stroke-dasharray:1] [&_path]:[stroke-dashoffset:1]",
                        inView && "[&_path]:animate-sketch-lines"
                    )}
                    width="100%"
                    fill="none"
                >
                    <path pathLength="1" d="M1500 72L220 72"></path>
                    <path pathLength="1" d="M1500 128L220 128"></path>
                    <path pathLength="1" d="M1500 189L220 189"></path>
                    <path pathLength="1" d="M220 777L220 1"></path>
                    <path pathLength="1" d="M538 777L538 128"></path>
                </svg>

                <CryptoCardLarge className={classNames(
                    "z-10 transition-opacity duration-1000",
                    inView ? "opacity-100" : "opacity-0"
                )} address="idx1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3206jxe" name="angelo.snr" type="Sonr" />
            </div>
        </div>
        </AspectRatio>
    );
};

export function CryptoIcon({ type }: { type: string }) {
    return (
        <div className="h-8 w-8 text-white/60">
            {getIcon(type)}
        </div>
    )
}

function getIcon(type: string) {
    if (type.includes("Ethereum")) {
        return <Icons.ethereum className="h-8 w-8 text-white/60" />
    } else if (type.includes("Bitcoin")) {
        return <Icons.bitcoin className="h-8 w-8 text-white/60" />
    } else if (type.includes("Sonr")) {
        return <Icons.sonr className="h-8 w-8 text-white/60" />
    }
    switch (type) {
        case "ETH":
            return <Icons.ethereum className="h-8 w-8 text-white/60" />
        case "BTC":
            return <Icons.bitcoin className="h-8 w-8 text-white/60" />
        case "SNR":
            return <Icons.sonr className="h-8 w-8 text-white/60" />
        default:
            return <Icons.ethereum className="h-8 w-8 text-white/60" />
    }
}
