import { cn } from "@/lib/utils";

export const Container = ({
    children,
    className,
}: {
    children: React.ReactNode;
    className?: string;
}) => {
    return (
        <div className={cn("mx-auto max-w-[120rem] px-8", className)}>
            {children}
        </div>
    );
};
