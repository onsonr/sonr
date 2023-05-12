import React, { useEffect, useRef } from 'react';

const colorPalettes: Record<string, string[]> = {
    sonr: [
        "#F68489",
        "#2BC0F6",
        "#805AEB",
        "#109CEF",
    ],
    palette1: [
        "#F78A9C",
        "#2EC6F7",
        "#8C62EB",
        "#11A8EE",
    ],
    palette2: [
        "#E75B72",
        "#1DB3D9",
        "#6E3BCA",
        "#0C80B3",
    ],
    palette3: [
        "#F49CA6",
        "#56C0E8",
        "#A37EF3",
        "#3DA1C9",
    ],
    palette4: [
        "#FFB6C1",
        "#8FD5F5",
        "#C2A0F4",
        "#4AC4E3",
    ],
    palette5: [
        "#D94C66",
        "#0FAAC5",
        "#5E2EB5",
        "#0092B2",
    ],
    bitcoin: [
        "#F7931A", // Vibrant orange
        "#FFB431", // Mellow yellow
        "#E56A00", // Darker orange
        "#FFCA7A", // Lighter mellow yellow
    ],

    ethereum: [
        "#3C3C3D", // Deep slate blue
        "#5A5A5C", // Slightly lighter slate blue
        "#1F1F20", // Darker slate blue
        "#828284", // Lighter slate blue
    ],

    usdc: [
        "#B22234", // Red
        "#KC3996B", // White
        "#3C3B6E", // Blue
        "#2BC0F6", // Lighter gray (alternative for white)
    ],
};

const AuroraBackground = ({ variant = "sonr", width = 700 }) => {
    const colors = colorPalettes[variant];
    return (<svg width={width} preserveAspectRatio="xMidYMid meet" viewBox="0 0 1920 1080" fill="none" xmlns="http://www.w3.org/2000/svg" className="rounded-3xl shadow-xl">
        <g clipPath="url(#clip0_42_44)">
            <rect width="1920" height="1080" fill="url(#paint0_radial_42_44)" />
            <g filter="url(#filter0_f_42_44)">
                <path
                    d="M1269.26 888.245C1269.26 1302.86 937.365 1638.97 527.949 1638.97C118.533 1638.97 -213.364 1302.86 -213.364 888.245C-213.364 473.63 118.533 137.518 527.949 137.518C937.365 137.518 1269.26 473.63 1269.26 888.245Z"
                    fill={colors[0]} />
                <path
                    d="M2538.96 1183.61C2538.96 1657.13 2177.05 2041 1730.61 2041C1284.17 2041 922.265 1657.13 922.265 1183.61C922.265 710.091 1284.17 326.225 1730.61 326.225C2177.05 326.225 2538.96 710.091 2538.96 1183.61Z"
                    fill={colors[1]} />
                <path
                    d="M1632.03 -42.9845C1632.03 423.74 1197.74 802.096 662.016 802.096C126.291 802.096 -308 423.74 -308 -42.9845C-308 -509.709 126.291 -888.065 662.016 -888.065C1197.74 -888.065 1632.03 -509.709 1632.03 -42.9845Z"
                    fill={colors[2]} />
                <path
                    d="M3091 -161.952C3091 370.477 2688.49 802.096 2191.96 802.096C1695.43 802.096 1292.92 370.477 1292.92 -161.952C1292.92 -694.381 1695.43 -1126 2191.96 -1126C2688.49 -1126 3091 -694.381 3091 -161.952Z"
                    fill={colors[3]} />
            </g>
        </g>
        <defs>
            <filter id="filter0_f_42_44" x="-753" y="-1571" width="4289" height="4057" filterUnits="userSpaceOnUse"
                colorInterpolationFilters="sRGB">
                <feFlood floodOpacity="0" result="BackgroundImageFix" />
                <feBlend mode="normal" in="SourceGraphic" in2="BackgroundImageFix" result="shape" />
                <feGaussianBlur stdDeviation="222.5" result="effect1_foregroundBlur_42_44" />
            </filter>
            <radialGradient id="paint0_radial_42_44" cx="0" cy="0" r="1" gradientUnits="userSpaceOnUse"
                gradientTransform="translate(359.599 577.241) rotate(7.41787) scale(906.584 560.025)">
                <stop stopColor="#E9CA9C" />
                <stop offset="1" stopColor="#0029FF" stopOpacity="0.33" />
            </radialGradient>
            <clipPath id="clip0_42_44">
                <rect width="1920" height="1080" fill="white" />
            </clipPath>
        </defs>
    </svg>
    )
}

export default AuroraBackground;
