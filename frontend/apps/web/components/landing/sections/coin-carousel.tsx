import React from 'react'
import icons from '../../../public/img/crypto/icons.json'
import { cn } from '@/lib/utils'

const CoinCarousel = () => {
    const rows = [[""], [""], [""]]
    const topRow = rows[0]
    const middleRow = rows[1]
    const bottomRow = rows[2]


    // Double the list of icons and distribute them into three rows
    const doubledIcons = [...icons, ...icons]
    doubledIcons.forEach((icon, i) => {
        rows[i % 3].push(icon)
    })

    return (
        <div className="mask-keyboard relative overflow-hidden pt-8">
            {/* Fade out effect */}
            <div className="absolute inset-y-0 left-0 w-40 bg-gradient-to-r from-black to-transparent"></div>
            <div className="absolute inset-y-0 right-0 w-40 bg-gradient-to-r from-transparent to-black"></div>

            <div className={cn("flex overflow-x-hidden whitespace-nowrap p-1.5")} key={"top"}>
                {topRow.map((icon, j) => (
                    <img src={icon} className="animate-carousel-slide-slow mx-2 inline-block w-[88px]" alt="" key={j} />
                ))}
            </div>
            <div className={cn("flex overflow-x-hidden whitespace-nowrap p-1.5")} key={"bottom"}>
                {middleRow.map((icon, j) => (
                    <img src={icon} className="animate-carousel-slide-fast mx-2 inline-block w-[88px]" alt="" key={j} />
                ))}
            </div>
            <div className={cn("flex overflow-x-hidden whitespace-nowrap p-1.5")} key={"bottom"}>
                {bottomRow.map((icon, j) => (
                    <img src={icon} className="animate-carousel-slide-medium mx-2 inline-block w-[88px]" alt="" key={j} />
                ))}
            </div>
        </div>
    )
}

export default CoinCarousel
