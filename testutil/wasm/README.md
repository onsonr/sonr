# WASM TEST

This file contains a simple script that uploads and inits a CW721 NFT contract.

To Run the Test simply update in the wallet.js

Run `yarn` to install the node modules

To Run The Test Run: `node wasm_test.js`

Successful run should print address in console(eg):

```shell
Upload succeeded. Receipt: {"originalSize":253184,"originalChecksum":"774e483a4a133ec4dc418aac96a233fff292e6efaf73afd78160de1b422ca3e3","compressedSize":83970,"compressedChecksum":"ed5f33ab36241d37435d566ce4c15a0afb9cf1f6e66cf4dcb1c
39629d1f8de87","codeId":1,"logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"/cosmwasm.wasm.v1.MsgStoreCode"},{"key":"module","value":"wasm"},{"key":"sender","value":"snr1vl4xg8quydtgj
9p7psqcpzs9x5azv534hdh49q"}]},{"type":"store_code","attributes":[{"key":"code_id","value":"1"}]}]}],"height":415,"transactionHash":"14A1981ADF3FEFB66254160D569C0B6545969BEF1440C2E8063C924E157529B6","gasWanted":0,"gasUsed":933124}


Contract instantiated at snr14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sxac83d
Contract Address snr14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sxac83d
snr14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sxac83d
```
