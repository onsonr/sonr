import { NextResponse } from 'next/server';
import { get } from '@vercel/edge-config';
import { NextApiRequest, NextApiResponse } from 'next';


export default async function handler(req: NextApiRequest, res: NextApiResponse) {
    const allowed = await get('allow_list');
    if (!allowed) {
        return NextResponse.error()
    }
    const { code } = req.query;
    for (const org of allowed as any) {
        if (code === org.code) {
            return res.status(200).json({ success: true, org: org })
        }
    }
    return res.status(200).json({ success: false })
}
