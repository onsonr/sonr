interface Organization {
    name?: string;
    domain?: string;
    code: string;
    allowed_emails: string[];
    blocked_emails?: string[];
    amount: number;
}

export type { Organization };
