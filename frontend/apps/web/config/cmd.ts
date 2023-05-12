interface CmdConfig {
    help: string
    cmds: {
        [key: string]: CmdItem[]
    }
}

interface CmdItem {
    name: string
    description: string
    type: 'filter' | 'form' | 'chat' | 'link'
    icon: string
    disabled?: boolean
    href?: string
}

export const cmdConfig: CmdConfig = {
    help: 'Filter Commands...',
    cmds: {
        'Accounts': [
            {
                name: 'Search Accounts...',
                description: 'Queries the blockchain for an alias, address, or DID.',
                type: 'filter',
                icon: '/icons/search.svg'
            },
            {
                name: 'Create New Account...',
                description: 'Create a new account on the blockchain.',
                type: 'form',
                icon: '/icons/new-account.svg'
            },
            {
                name: 'Link New Device...',
                description: 'Link a new device to your account.',
                type: 'link',
                icon: '/icons/broadcast-tx.svg'
            },
            {
                name: 'Open Inbox for Account...',
                description: 'Send a message to another account.',
                type: 'chat',
                icon: '/icons/open-inbox.svg'
            },
        ],
        'Services': [
            {
                name: 'Search Records...',
                description: 'Queries the blockchain for a service record.',
                type: 'filter',
                icon: '/icons/search.svg'
            },
            {
                name: 'Register New Record...',
                description: 'Create a new service record on the blockchain.',
                type: 'form',
                icon: '/icons/new-service.svg'
            },
            {
                name: 'Edit Record...',
                description: 'Edit an existing service record.',
                type: 'form',
                disabled: true,
                icon: '/icons/edit-service.svg'
            },
        ],
        'Help': [
            {
                name: 'Documentation',
                description: 'Read the documentation for Sonr.',
                type: 'link',
                href: 'https://snr.la/doc',
                icon: '/icons/visit-docs.svg'
            },
            {
                name: 'Contact Support',
                description: 'Contact the dev team for help.',
                type: 'link',
                href: 'mailto:team@sonr.network',
                icon: '/icons/support.svg'
            }
        ]
    }

}
