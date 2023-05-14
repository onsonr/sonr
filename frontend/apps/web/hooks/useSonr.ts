"use client"

import * as React from 'react';
import { useUser } from './useUser';

import { SonrClient } from '../../../packages/client/lib';

import { DidDocument, QueryAliasResponse } from '../../../packages/client/lib/types';
import { User } from '../../../packages/client/lib/types/user';
import { SonrLoginProps, SonrRegisterProps } from '../../../packages/client/lib/types/props';



export interface Sonr {
    checkAlias: (alias: string) => Promise<QueryAliasResponse>;
    login: ({ alias, onCredentialSet, onLoginComplete }: SonrLoginProps) => Promise<void>;
    register: ({ alias, onCredentialSet, onRegisterComplete }: SonrRegisterProps) => Promise<void>;
    user: User | undefined;
    logout: () => void;
    client: SonrClient;
    didDocument?: DidDocument;
}

export const useSonr = (): Sonr => {
    let origin = "sonr.id"
    if (process.env.NODE_ENV === "development") {
        origin = "localhost"
    }

    const [user, setUser] = useUser();
    const [didDocument, setDidDocument] = React.useState<DidDocument | undefined>(undefined);
    const client: SonrClient = new SonrClient(origin);

    const checkAlias = async (alias: string) => {
        const resp = await client.did.getByAlias(alias);
        return resp;
    };

    const login = async ({alias, onLoginComplete, onCredentialSet}: SonrLoginProps) => {
        const resp = await client.login({ alias, onLoginComplete, onCredentialSet });
        setDidDocument(resp.did_document);
        let newUsr: User = {
            did: resp.did_document.id,
            didDocument: resp.did_document,
            username: alias,
            address: resp.address,
        };
        setUser(newUsr);
    };

    const register = async ({ alias, onCredentialSet, onRegisterComplete: onRegisterComplete }: SonrRegisterProps) => {
        const resp = await client.register({ alias, onCredentialSet, onRegisterComplete });
        setDidDocument(resp.primary);
        let newUsr: User = {
            did: resp.primary.id,
            didDocument: resp.primary,
            username: alias,
            address: resp.address,
        };
        setUser(newUsr);
    };

    const logout = () => {
        localStorage.removeItem('user');
    };

    return {
        checkAlias,
        login,
        register,
        user,
        logout,
        client,
        didDocument,
    };
};
