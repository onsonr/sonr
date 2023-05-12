import kv from '@vercel/kv';
import { NextApiRequest, NextApiResponse } from 'next';

export default async function handler(
    request: NextApiRequest,
    response: NextApiResponse,
) {
    kv.sadd('users', 'me');
    const user = await kv.hgetall('user:me');
    return response.status(200).json(user);
}
