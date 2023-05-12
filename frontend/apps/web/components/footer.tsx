import Link from "next/link";
import { Container } from "./container";
import { Icons } from "@/components/icons";


const footerLinks = [
    {
        title: "Technology",
        links: [
            { title: "Features", href: "#features" },
            { title: "Changelog", href: "https://changelog.sonr.io" },
            { title: "Docs", href: "https://snr.la/doc" },
            { title: "Roadmap", href: "https://roadmap.sonr.io" },
        ],
    },
    {
        title: "Company",
        links: [
            { title: "About us", href: "https://sonr.io/about" },
            { title: "Blog", href: "https://blog.sonr.io" },
            { title: "Careers", href: "#" },
            { title: "Brand", href: "https://design.sonr.io" },
        ],
    },
    {
        title: "Resources",
        links: [
            { title: "Community", href: "https://snr.la/dev-chat" },
            { title: "Contact", href: "https://pradn.me/telegram" },
            { title: "Investors", href: "https://snr.la/seed-vc" },
            { title: "Token Economics", href: "https://snr.la/tokenomics" },
        ],
    },
    {
        title: "Developers",
        links: [
            { title: "API", href: "https://api.sonr.ws" },
            { title: "Status", href: "https://status.sonr.io" },
            { title: "GitHub", href: "https://snr.la/gh" },
        ],
    },
];

export const Footer = () => (
    <footer className="border-transparent-white mt-12 border-t py-[5.6rem] text-sm">
        <Container className="flex flex-col justify-between lg:flex-row">
            <div>
                <div className="flex h-full flex-row justify-between lg:flex-col">
                    <div className="text-grey mr-3 flex items-center">
                        <Icons.sonr className="text-off-white/70 mr-2 h-7 w-7" /> <span className="pr-1.5 font-semibold">Sonr </span>  - The Internet Rebuilt for you
                    </div>
                    <div className="text-grey mt-auto flex items-center space-x-5">
                        <div>

                        </div>
                        <Icons.twitter className="w-6 cursor-pointer hover:text-white/80" onClick={(e) => {
                            e.preventDefault();
                            window.open("https://snr.la/tw");
                        }} />
                        <Icons.gitHub className="w-6 cursor-pointer hover:text-white/80" onClick={(e) => {
                            e.preventDefault();
                            window.open("https://snr.la/gh");
                        }} />
                        <Icons.discord className="w-6 cursor-pointer hover:text-white/80" onClick={(e) => {
                            e.preventDefault();
                            window.open("https://snr.la/dev-chat");
                        }} />
                    </div>
                </div>
            </div>
            <div className="flex flex-wrap">
                {footerLinks.map((column) => (
                    <div
                        key={column.title}
                        className="mt-10 min-w-[50%] lg:mt-0 lg:min-w-[18rem]"
                    >
                        <h3 className="mb-3 font-medium">{column.title}</h3>
                        <ul>
                            {column.links.map((link) => (
                                <li key={link.title} className="[&_a]:last:mb-0">
                                    <Link
                                        className="text-grey hover:text-off-white mb-3 block transition-colors"
                                        href={link.href}
                                    >
                                        {link.title}
                                    </Link>
                                </li>
                            ))}
                        </ul>
                    </div>
                ))}
            </div>
        </Container>
    </footer>
);
