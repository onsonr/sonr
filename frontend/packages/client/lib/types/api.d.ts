import { DidDocument } from "./did";
import { Account } from "./user";
interface GetAccountResponse {
    success: boolean;
    account: Account;
    coin_type: string;
    name: string;
    address: string;
}
interface ListAccountsResponse {
    success: boolean;
    accounts: Account[];
}
interface CreateAccountResponse {
    success: boolean;
    coin_type: string;
    did_document: DidDocument;
    new_account: Account;
}
interface QueryDocumentResponse {
    success: boolean;
    account_address: string;
    did_document: DidDocument;
}
interface ListDocumentsResponse {
    documents: DidDocument[];
    count: number;
}
interface QueryAttestionResponse {
    alias: string;
    origin: string;
    attestion_options: string;
    challenge: string;
    ucw_id: string;
}
interface RegistrationResponse {
    success: boolean;
    address: string;
    did: string;
    primary: DidDocument;
    accounts: Account[];
    tx_hash: string;
    jwt: string;
}
interface LoginRequest {
    account_address: string;
    username: string;
    credential_response: string;
    origin: string;
}
interface LoginResponse {
    success: boolean;
    address: string;
    did: string;
    keyshares: Record<string, Uint8Array>;
    did_document: DidDocument;
    jwt: string;
}
interface QueryAliasResponse {
    available: boolean;
    alias: string;
    did: string;
    document: DidDocument;
    address: string;
}
interface QueryAssertionResponse {
    success: boolean;
    origin: string;
    assertion_options: string;
    address: string;
}
interface SendMessageResponse {
    success: boolean;
    to: string;
    from: string;
    message: string;
}
interface BlockResponse {
    jsonrpc: string;
    id: number;
    result: {
        block_id: {
            hash: string;
            parts: {
                total: number;
                hash: string;
            };
        };
        block: {
            header: {
                version: {
                    block: string;
                };
                chain_id: string;
                height: string;
                time: string;
                last_block_id: {
                    hash: string;
                    parts: {
                        total: number;
                        hash: string;
                    };
                };
                last_commit_hash: string;
                data_hash: string;
                validators_hash: string;
                next_validators_hash: string;
                consensus_hash: string;
                app_hash: string;
                last_results_hash: string;
                evidence_hash: string;
                proposer_address: string;
            };
            data: {
                txs: any[];
            };
            evidence: {
                evidence: any[];
            };
            last_commit: {
                height: string;
                round: number;
                block_id: {
                    hash: string;
                    parts: {
                        total: number;
                        hash: string;
                    };
                };
                signatures: {
                    block_id_flag: number;
                    validator_address: string;
                    timestamp: string;
                    signature: string;
                }[];
            };
        };
    };
}
interface AccountInfo {
    address: string;
    name: string;
    did: string;
    coinType: string;
    chainId: string;
    publicKey: string;
    type: string;
}
interface ListAccountsResult {
    success: boolean;
    accounts: {
        [coinType: string]: AccountInfo[];
    };
}
export type { GetAccountResponse, ListAccountsResponse, CreateAccountResponse, QueryDocumentResponse, ListDocumentsResponse, QueryAttestionResponse, RegistrationResponse, LoginRequest, LoginResponse, QueryAliasResponse, QueryAssertionResponse, SendMessageResponse, BlockResponse, ListAccountsResult, AccountInfo, };
