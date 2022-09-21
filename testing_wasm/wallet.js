//import message from "@cosmostation/cosmosjs/src/messages/proto.js";
import fs from "fs";
import fetch from "node-fetch";
//import {Cosmos} from "@cosmostation/cosmosjs";
import {SigningCosmWasmClient} from "@cosmjs/cosmwasm-stargate";
import {DirectSecp256k1HdWallet} from "@cosmjs/proto-signing";
import {calculateFee, GasPrice} from "@cosmjs/stargate";
import wasmTxType from "cosmjs-types/cosmwasm/wasm/v1/tx.js";
import {toUtf8} from "@cosmjs/encoding";

const {MsgExecuteContract, MsgSend} = wasmTxType;

const debug = true

const chainId = "sonr"
//const lcdUrl = "http://localhost:1317"
const endpoint = "http://localhost:26657";


// Copy Memonic from the Terminal and UPDATE IT HERE
export const mnemonic = "service ball cushion genius salmon cross find grape inject long inquiry rent liberty object raw nest alarm category crime adapt awesome taxi visual exhibit"


export class Wallet {
    wallet_address;
    publicKey;
    privateKey;
    client;
    gasPrice;
    memonic;

    constructor(memonic) {

        this.memonic = memonic;
    }

    async initialize() {
        const wallet = await DirectSecp256k1HdWallet.fromMnemonic(this.memonic, {prefix: "snr"});
        const account = await wallet.getAccounts();
        this.wallet_address = account[0].address;
        this.client = await SigningCosmWasmClient.connectWithSigner(endpoint, wallet, {gasPrice: GasPrice.fromString("1snr")});
    }

    async sign_and_broadcast(messages) {
        const memo = "sign_and_broadcast_memo";
        console.log(messages)
        return this.client.signAndBroadcast(this.wallet_address, messages, "auto", memo)
    }

    async send_funds(to_address, amount, denom) {

        return this.sign_and_broadcast([{
            typeUrl: "/cosmos.bank.v1beta1.MsgSend",
            value: {
                fromAddress: this.wallet_address,
                toAddress: to_address,
                amount: [{amount: amount, denom: denom}]
            }
        }
        ])
    }

    async execute_contract(msg, contractAddress, coins) {
        let msg_list = []
        if (Array.isArray(msg)) {
            msg.forEach((msg) => {
                msg_list.push(this.get_execute(msg, contractAddress, coins))
            })

        } else {
            msg_list = [
                this.get_execute(msg, contractAddress, coins)
            ]
        }
        console.log("execute_contract is called")
        console.log(JSON.stringify(msg_list))
        let response = await this.sign_and_broadcast(msg_list)
        console.log(response)
        return response
    }

    get_execute(msg, contract, coins) {

        if (typeof coins === "object") {
            coins = [coins]
        } else {
            coins = []
        }

        const executeContractMsg = {
            typeUrl: "/cosmwasm.wasm.v1.MsgExecuteContract",
            value: MsgExecuteContract.fromPartial({
                sender: this.wallet_address,
                contract: contract,
                msg: (0, toUtf8)(JSON.stringify(msg)),
                funds: coins,
            }),
        };
        return executeContractMsg;
    }

    query(address, query) {

        return this.client.queryContractSmart(address, JSON.stringify(query))
    }


    async upload(file) {
        const code = fs.readFileSync(file);
        const uploadReceipt = await this.client.upload(
            this.wallet_address,
            code,
            "auto",
            "Uploading contract",
        );
        console.info(`Upload succeeded. Receipt: ${JSON.stringify(uploadReceipt)}`);
        return uploadReceipt
    }

    async init(code_id, contract_init) {

        const instantiateFee = calculateFee(500, GasPrice.fromString("0.0001snr"));
        const {contractAddress} = await this.client.instantiate(
            this.wallet_address,
            code_id,
            contract_init,
            "some_label",
            "auto",
            {
                memo: `Create a instance of contract`,
                admin: this.wallet_address,
            },
        );
        console.info(`Contract instantiated at ${contractAddress}`);
        return contractAddress

    }


    sleep(time) {
        return new Promise((resolve) => setTimeout(resolve, time));
    }

    queryBankUusd(address) {
        let api = "/cosmos/bank/1beta1/balances/";
        return fetch(this.url + api + address).then(response => response.json())
    }


}


let wallet = new Wallet(mnemonic)
await wallet.initialize();
