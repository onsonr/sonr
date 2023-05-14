import { User } from '../../../packages/client/lib/types/user';
import * as React from 'react';

type UseUser = [
    User | undefined,
    (account: User | undefined) => void
];

export const useUser = (): UseUser => {
    const [user, setUser] = React.useState<User | undefined>(undefined);

    React.useEffect(() => {
        if (user !== undefined) {
            localStorage.setItem('user', JSON.stringify(user));
        }
    }, [user]);

    React.useEffect(() => {
        const userBz = localStorage.getItem('user');
        if (userBz !== null) {
            let parsedUser = JSON.parse(userBz);
            if (parsedUser.didDocument !== null && parsedUser.didDocument !== undefined) {
                setUser(parsedUser);
            } else {
                let isDev = window.location.hostname === 'localhost';
                if (isDev) {
                    localStorage.removeItem('user');
                }
            }
        }
    }, []);

    return [user, setUser];
};
