import { NextApiRequest, NextApiResponse } from 'next';
import { SonrNodeResponse } from '@/types/node';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
    const validators = [
        {
            name: 'ashburn',
            rpc: 'https://rpc.ashburn.sonr.zone/status',
        },
        {
            name: 'brooklyn',
            rpc: 'https://rpc.brooklyn.sonr.zone/status',
        },
        {
            name: 'ibiza',
            rpc: 'https://rpc.ibiza.sonr.zone/status',
        },
        {
            name: 'rome',
            rpc: 'https://rpc.rome.sonr.zone/status',
        },
    ];

    try {
        const responses = await Promise.all(
            validators.map(async (validator) => {
                const response = await fetch(validator.rpc);
                const data: SonrNodeResponse = await response.json();
                return { [validator.name]: data };
            })
        );
        const result = Object.assign({}, ...responses);
        return res.status(200).json(result);
    } catch (error) {
        return res.status(500).json({ error: 'Failed to fetch validator status' });
    }
}
