import {NftContractAddress} from "./constants.js";
import {Wallet} from "./wallet.js";

let test_name = "Test"

export class NftTestEngine extends Wallet {

    constructor(
        mnemonic,
        minter,
        name,
        symbol,
    ) {
        super(mnemonic);
        this.contract_path = NftContractAddress
        this.init_message = {
            "minter": minter,
            "name": name,
            "symbol": symbol,
        }

    }

    async setup() {
        await this.initialize();
        this.code_id = await this.upload(this.contract_path)
        console.log()
        this.contract_address = await this.init(this.code_id.codeId, this.init_message);
        console.log(`Contract Address ${this.contract_address}`)

    }

    async transfer_nft(
        recipient,
        token_id,
    ) {
        console.log("Executing transfer_nft")
        let message = {
            transfer_nft: {

                "recipient": recipient,

                "token_id": token_id,


            }
        }
        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async send_nft(
        contract,
        msg,
        token_id,
    ) {
        console.log("Executing send_nft")
        let message = {
            send_nft: {

                "contract": contract,

                "msg": msg,

                "token_id": token_id,


            }
        }
        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async approve(
        spender,
        token_id,
    ) {
        console.log("Executing approve")
        let message = {
            approve: {

                "spender": spender,

                "token_id": token_id,


            }
        }
        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async revoke(
        spender,
        token_id,
    ) {
        console.log("Executing revoke")
        let message = {
            revoke: {

                "spender": spender,

                "token_id": token_id,


            }
        }
        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async approve_all(
        operator,
    ) {
        console.log("Executing approve_all")
        let message = {
            approve_all: {

                "operator": operator,


            }
        }
        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async revoke_all(
        operator,
    ) {
        console.log("Executing revoke_all")
        let message = {
            revoke_all: {

                "operator": operator,


            }
        }
        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async mint() {
        console.log("Executing mint")
        let message = {
            mint: {}
        }
        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async burn(
        token_id,
    ) {
        console.log("Executing burn")
        let message = {
            burn: {

                "token_id": token_id,


            }
        }
        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async owner_of(
        token_id,
    ) {
        console.log("Querying owner_of")
        let message = {
            owner_of: {

                "token_id": token_id,


            }
        }


        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async approval(
        spender,
        token_id,
    ) {
        console.log("Querying approval")
        let message = {
            approval: {

                "spender": spender,

                "token_id": token_id,


            }
        }


        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async approvals(
        token_id,
    ) {
        console.log("Querying approvals")
        let message = {
            approvals: {

                "token_id": token_id,


            }
        }


        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async all_operators(
        owner,
    ) {
        console.log("Querying all_operators")
        let message = {
            all_operators: {

                "owner": owner,


            }
        }


        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async num_tokens() {
        console.log("Querying num_tokens")
        let message = {
            num_tokens: {}
        }


        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async contract_info() {
        console.log("Querying contract_info")
        let message = {
            contract_info: {}
        }


        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async nft_info(
        token_id,
    ) {
        console.log("Querying nft_info")
        let message = {
            nft_info: {

                "token_id": token_id,


            }
        }


        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async all_nft_info(
        token_id,
    ) {
        console.log("Querying all_nft_info")
        let message = {
            all_nft_info: {

                "token_id": token_id,


            }
        }


        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async tokens(
        owner,
    ) {
        console.log("Querying tokens")
        let message = {
            tokens: {

                "owner": owner,


            }
        }


        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async all_tokens() {
        console.log("Querying all_tokens")
        let message = {
            all_tokens: {}
        }


        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


    async minter() {
        console.log("Querying minter")
        let message = {
            minter: {}
        }


        let response = this.execute_contract(message, this.contract_address)
        console.log(response)
        return response


    }


}