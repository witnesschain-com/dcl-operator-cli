# DCL Operator CLI

## Description
dcl-operator is a command-line interface (CLI) tool for registering provers and challengers.

## Installation
Installation instructions for building from source is available in 
[docs/install.md](docs/install.md).

## Register Provers

Create a `prover.json` file to provide provers and operator keys.

Sample of prover config file:-

```json
{
  "prover_private_keys": [
    "<raw-prover-private-key e.g. 01234567890abcdef01234567890abcdef01234567890abcdef01234567890abcdef01234567890abcdef>",
    "<raw-prover-private-key e.g. 01234567890abcdef01234567890abcdef01234567890abcdef01234567890abcdef01234567890abcdef>"
  ],
  "operator_private_key": "<operator private key e.g. 01234567890abcdef01234567890abcdef01234567890abcdef01234567890abcdef01234567890abcdef>",
  "eth_rpc_url": "https://blue-orangutan-blockscout.eu-north-2.gateway.fm/",
}
```

Run the following command to register the provers:-

```bash
$ dcl-operator registerProver --config-file prover.json
```

You will get similar output on running above commands,
```bash
Using config file path : prover.json
b961628baffcd195edb09542c8c1eb3b70ba98efa0824f1382edc624a0de1600
Connection successful :  1237146866
github.com/witnes .. │ Oct 25 15:46:33 2024 │ ➤ keystore: raw://0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
proverAddress: 0x9fa70F82F6A872e5a0324eeF3A824Ec22A5D61Fb
github.com/witnes .. │ Oct 25 15:46:33 2024 │ ➤ keystore: raw://0x9fa70F82F6A872e5a0324eeF3A824Ec22A5D61Fb
Tx sent: https://blue-orangutan-blockscout.eu-north-2.gateway.fm/tx/0x39b65081231e0918a4b812892cb1332c5663f67963c1e396e104ae2a83e6da65
Transaction executed successfully, logs are ...
[0xc000338000]
```

For mainnet, please replace `eth_rpc_url` with 
`https://rpc.witnesschain.com`.


## Register Challengers

Similarly, create a `challenger.json` file to provide challengers and operator keys.

```json
{
  "challenger_private_keys": [
    "<raw-challenger-private-key e.g. 01234567890abcdef01234567890abcdef01234567890abcdef01234567890abcdef01234567890abcdef>",
    "<raw-challenger-private-key e.g. 01234567890abcdef01234567890abcdef01234567890abcdef01234567890abcdef01234567890abcdef>"
  ],
  "operator_private_key": "<raw-operator-private-key>",
  "eth_rpc_url": "https://blue-orangutan-rpc.eu-north-2.gateway.fm",
}
```

Run the following command to register the challengers:-

```bash
$ dcl-operator registerChallenger --config-file challenger.json
```

You get output similar to the following:

```bash
Using config file path : challenger.json
ae45757881825d0c9a890d0b819d88f9048a850eb3aed11def164655aece975e
Connection successful :  1237146866
github.com/witnes .. │ Oct 25 15:44:34 2024 │ ➤ keystore: raw://0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
challengerAddress: 0x95e46074F96c6DEa2bE91BE9678Eb100A3826310
github.com/witnes .. │ Oct 25 15:44:34 2024 │ ➤ keystore: raw://0x95e46074F96c6DEa2bE91BE9678Eb100A3826310
Tx sent: https://blue-orangutan-blockscout.eu-north-2.gateway.fm/tx/0x9b42f917022c7daa3ded2f98f01d2661f61e29fbf72cbbf7416e5eb92afed587
Transaction executed successfully, logs are ...
[0xc0002a22c0]
```

For mainnet, please replace `eth_rpc_url` with 
`https://rpc.witnesschain.com`.

Web3secret keys are also supported, if you wish to use web3secret 
format, please visit [docs/web3secret.md](docs/web3secret.md).
