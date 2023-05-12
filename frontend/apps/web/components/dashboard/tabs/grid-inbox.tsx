"use client"

import { Card, CardHeader, CardFooter, CardTitle, CardContent, CardDescription } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import { Button } from "@/components/ui/button";
import { ScrollArea } from "../../ui/scroll-area";
import { cn } from "@/lib/utils";
import { Input } from "../../ui/input";
import { Icons } from "../../icons";
import { Avatar, AvatarImage } from "@radix-ui/react-avatar";
import { AvatarFallback } from "../../ui/avatar";
import { Tabs } from "@/components/ui/tabs";


export default function GridInbox() {
    return (
        <Card className="p-8">
            <CardHeader>
                <CardTitle>Messages</CardTitle>
            </CardHeader>
            <div className="grid h-max md:grid-cols-2 lg:grid-cols-10">
                <Card className="col-span-4 h-full rounded-r-none border-r-0">
                    <CardHeader>
                        <CardTitle>Conversations</CardTitle>
                    </CardHeader>
                    <CardContent className="mb-8 p-0">
                        <ScrollArea className="h-96">
                            <Card className="flex gap-x-2.5 rounded-none p-3">
                                <Avatar>
                                    <AvatarImage src="/favicon.ico"></AvatarImage>
                                </Avatar>
                                <div>
                                    <CardTitle className="text-sm">1A32qrF213FsEfg</CardTitle>
                                    <CardDescription className="text-xs">
                                        These cookies are essential in order to use the website and use its features.
                                    </CardDescription>
                                </div>
                            </Card>
                            <Card className="flex gap-x-2.5 rounded-none p-3">
                                <Avatar>
                                    <AvatarImage src="/favicon.ico"></AvatarImage>
                                </Avatar>
                                <div>
                                    <CardTitle className="text-sm">1A32qrF213FsEfg</CardTitle>
                                    <CardDescription className="text-xs">
                                        These cookies are essential in order to use the website and use its features.
                                    </CardDescription>
                                </div>
                            </Card>
                            <Card className="flex gap-x-2.5 rounded-none p-3">
                                <Avatar>
                                    <AvatarImage src="/favicon.ico"></AvatarImage>
                                </Avatar>
                                <div>
                                    <CardTitle className="text-sm">1A32qrF213FsEfg</CardTitle>
                                    <CardDescription className="text-xs">
                                        These cookies are essential in order to use the website and use its features.
                                    </CardDescription>
                                </div>
                            </Card>
                            <Card className="flex gap-x-2.5 rounded-none p-3">
                                <Avatar>
                                    <AvatarImage src="/favicon.ico"></AvatarImage>
                                </Avatar>
                                <div>
                                    <CardTitle className="text-sm">1A32qrF213FsEfg</CardTitle>
                                    <CardDescription className="text-xs">
                                        These cookies are essential in order to use the website and use its features.
                                    </CardDescription>
                                </div>
                            </Card>
                            <Card className="flex gap-x-2.5 rounded-none p-3">
                                <Avatar>
                                    <AvatarImage src="/favicon.ico"></AvatarImage>
                                </Avatar>
                                <div>
                                    <CardTitle className="text-sm">1A32qrF213FsEfg</CardTitle>
                                    <CardDescription className="text-xs">
                                        These cookies are essential in order to use the website and use its features.
                                    </CardDescription>
                                </div>
                            </Card>
                            <Card className="flex gap-x-2.5 rounded-none p-3">
                                <Avatar>
                                    <AvatarImage src="/favicon.ico"></AvatarImage>
                                </Avatar>
                                <div>
                                    <CardTitle className="text-sm">1A32qrF213FsEfg</CardTitle>
                                    <CardDescription className="text-xs">
                                        These cookies are essential in order to use the website and use its features.
                                    </CardDescription>
                                </div>
                            </Card>
                            <Card className="flex gap-x-2.5 rounded-none p-3">
                                <Avatar>
                                    <AvatarImage src="/favicon.ico"></AvatarImage>
                                </Avatar>
                                <div>
                                    <CardTitle className="text-sm">1A32qrF213FsEfg</CardTitle>
                                    <CardDescription className="text-xs">
                                        These cookies are essential in order to use the website and use its features.
                                    </CardDescription>
                                </div>
                            </Card>
                            <Card className="flex gap-x-2.5 rounded-none p-3">
                                <Avatar>
                                    <AvatarImage src="/favicon.ico"></AvatarImage>
                                </Avatar>
                                <div>
                                    <CardTitle className="text-sm">1A32qrF213FsEfg</CardTitle>
                                    <CardDescription className="text-xs">
                                        These cookies are essential in order to use the website and use its features.
                                    </CardDescription>
                                </div>
                            </Card>
                            <Card className="flex gap-x-2.5 rounded-none p-3">
                                <Avatar>
                                    <AvatarImage src="/favicon.ico"></AvatarImage>
                                </Avatar>
                                <div>
                                    <CardTitle className="text-sm">1A32qrF213FsEfg</CardTitle>
                                    <CardDescription className="text-xs">
                                        These cookies are essential in order to use the website and use its features.
                                    </CardDescription>
                                </div>
                            </Card>
                            <Card className="flex gap-x-2.5 rounded-none p-3">
                                <Avatar>
                                    <AvatarImage src="/favicon.ico"></AvatarImage>
                                </Avatar>
                                <div>
                                    <CardTitle className="text-sm">1A32qrF213FsEfg</CardTitle>
                                    <CardDescription className="text-xs">
                                        These cookies are essential in order to use the website and use its features.
                                    </CardDescription>
                                </div>
                            </Card>
                            <Card className="flex gap-x-2.5 rounded-none p-3">
                                <Avatar>
                                    <AvatarImage src="/favicon.ico"></AvatarImage>
                                </Avatar>
                                <div>
                                    <CardTitle className="text-sm">1A32qrF213FsEfg</CardTitle>
                                    <CardDescription className="text-xs">
                                        These cookies are essential in order to use the website and use its features.
                                    </CardDescription>
                                </div>
                            </Card>
                            <Card className="flex gap-x-2.5 rounded-none p-3">
                                <Avatar>
                                    <AvatarImage src="/favicon.ico"></AvatarImage>
                                </Avatar>
                                <div>
                                    <CardTitle className="text-sm">1A32qrF213FsEfg</CardTitle>
                                    <CardDescription className="text-xs">
                                        These cookies are essential in order to use the website and use its features.
                                    </CardDescription>
                                </div>
                            </Card>
                        </ScrollArea>
                    </CardContent>
                    <CardFooter>
                        <Tabs>

                        </Tabs>
                    </CardFooter>
                </Card>

                <Card className="border-l-none col-span-6 rounded-l-none">
                    <CardHeader>
                        <CardTitle>0x121rasdflnas9203u</CardTitle>
                        <CardDescription>
                            You made 265 authentications this month.
                        </CardDescription>
                    </CardHeader>
                    <CardContent>
                        <ScrollArea className="h-96">
                            <div className="grid grid-cols-7">
                                <ChatMessage message="HEllo" sender="Me" position="left" />
                                <ChatMessage message="HEllo" sender="Me" position="left" disableSender />
                                <ChatMessage message="HEllo" sender="Me" position="right" />
                                <ChatMessage message="HEllo" sender="Me" position="left" />
                            </div>
                        </ScrollArea>
                    </CardContent>
                    <CardFooter className="gap-x-2">
                        <Input>
                        </Input>
                        <Button className="bg-blue-700">
                            <Icons.sendMsg className="w-4 fill-white text-white" />
                        </Button>
                    </CardFooter>
                </Card>
            </div>
        </Card>
    )
}

interface ChatMessageProps {
    message: string
    sender: string
    position: 'left' | 'right'
    disableSender?: boolean
}

export function ChatMessage({ message, sender, position, disableSender = false }: ChatMessageProps) {
    return position == 'left' ? (
        <div className="col-start-1 col-end-4 m-2">
            {!disableSender && (
                <div className="chat-header pb-1 text-sm font-semibold">
                    {sender}
                </div>
            )}
            <div className="rounded-3xl bg-slate-700 px-2 py-3"><p className="text-md">{message}</p></div>
        </div>
    ) : (
        <div className="col-start-4 col-end-7 m-2">
            {!disableSender && (
                <div className="chat-header pb-1 pl-2 text-right text-sm font-semibold">
                    {sender}
                </div>
            )}
                <div className="bg-sonr rounded-3xl px-2 py-3 text-black"><p className="text-md">{message}</p></div>
        </div>
    )
}

interface ConversationItemProps {
    address: string
    name: string
    message: string
    avatar: string
}

export function ConversationItem({ address, name, message, avatar }: ConversationItemProps) {
    return (
        <Card className="flex gap-x-2.5 rounded-none p-4">
            <Avatar>
                <AvatarImage src={avatar}></AvatarImage>
                <AvatarFallback>{name.at(0)}{name.at(1)}</AvatarFallback>
            </Avatar>
            <div>
                <CardTitle className="text-sm">{address}</CardTitle>
                <CardDescription className="text-xs">
                    {message}
                </CardDescription>
            </div>
        </Card>
    )
}
