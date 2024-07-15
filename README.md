# DCL Operator CLI

## Description
dcl-operator is a command-line interface (CLI) tool for registering provers and challengers.
You can refer to the [How to use the config files](#how-to-use-the_config-files) section to understand how to use the config files.

## Installation
You can get the dcl-operator cli by building from source

### Pre-requisite

gocryptfs should be installed on the machine you will use the CLI tool on. The installation and running of gocryptfs is easier and smoother on Unbuntu/Debian systems. So, to use the CLI, those would be the ideal systems

1. **Building from source**
   - clone the repository 
    ```
    git clone https://github.com/witnesschain-com/dcl-operator-cli.git
    ```

   - Build the binaries
    ```
    cd operator-cli/dcl-operator-cli
    ./build
    ```


## How to use
Once you have the dcl-operator installed, you can directly use the exectable -

```
dcl-operator command [command options] [arguments...]
```
Note: In case you haven't exported the path for dcl-operator executable, you can start the cli by `./dcl-operator`

## Commands available
| Command | Description |
|----------|----------|
|keys | Used to store private keys in an encrypted format |
|registerProver | Used to register prover |
|deRegisterProver | Used to deregister prover |
|registerChallenger | Used to register challenger |
|deRegisterChallenger| Used to deregister challenger |

## How to use the config files
The structure and details in the config file might differ based on the functionality you are using. Enter the following details in the config file. Changing the key names in the json file will lead to misbehavior

### For the challenger

#### related commands - registerChallenger, deRegisterChallenger
Default file: operator-challenger-config.json.template (reference file)

| Field | Description |
|----------|----------|
|challenger_private_keys | Private keys (raw), key names or paths of the challenger keys |
|operator_private_key | Private key of the operator(on which the actions will be performed) |
|challenger_registry_address | Address of the ChallengerRegistry contract(used to register and deregister challenger) |
|eth_rpc_url | The RPC URL where you want to perform the transactions |
|chain_id | Chain ID of the respective chain |
|gas_limit | The gas limit you want to set while sending the transactions |
|tx_receipt_timeout| Timeout in seconds for waiting of tx receipts |
|expiry| Expiry in days after which the challenger signature becomes invalid |
|use_encrypted_keys| Indicates if you want to store and use the private keys in encrypted format |

### For the prover

#### related commands - registerProver, deRegisterProver
Default file: operator-prover-config.json.template (reference file)

| Field | Description |
|----------|----------|
|prover_private_keys | Private keys (raw), key names or paths of the prover keys |
|operator_private_key | Private key of the operator(on which the actions will be performed) |
|prover_registry_address | Address of the ProverRegistry contract(used to register and deregister prover) |
|eth_rpc_url | The RPC URL where you want to perform the transactions |
|chain_id | Chain ID of the respective chain |
|gas_limit | The gas limit you want to set while sending the transactions |
|tx_receipt_timeout| Timeout in seconds for waiting of tx receipts |
|expiry| Expiry in days after which the prover signature becomes invalid |
|use_encrypted_keys| Indicates if you want to store and use the private keys in encrypted format |

## How to use the encrypted keys

To store and use the keys in an encrypted format, use the `use_encrypted_keys` field in the config. If it is set to `true`, the CLI tool will use the keys stored in the encrypted format in the path given in `keys_directory_path/.encrypted_keys` or the default location, `~/.witnesschain/cli/.encrypted_keys`. The CLI uses gocrytpfs to achieve this encryption. Read more about it [here](https://github.com/rfjakob/gocryptfs)

### Sub-commands
The `keys` command has 4 sub-commands. Let's look at them one by one -
1. `init` - This command is used to initiate a new gocryptfs file system. This will ask you to input a password which will be required whenever we need to access these file. It also performs password validations for stromger passwords. To bypass this validation in the test environment, there is a flag `--insecure(-i)`. Once the command is successfully run, all other actions will need this password.
```
$ dcl-operator keys init

Creating directory:  .encrypted_keys
Creating directory:  .decrypted_keys
Enter password to init: **********
```
After this command, two directories .encrypted_keys and .decrypted_keys are created. The names indicate their functions. Once this is done, we don't need to do it again, unless the .encrypted_keys or .decrypted_keys are tampered with.

2. `create` - This command will create a new key. This command will need a flag --key-name(-k). This will be the name of the key which will be referred in the future by the CLI. This name should be mentioned in the config file to extract the corresponding private key. When you run this command, it will ask for password to mount and then ask you to enter the private key
```
$ dcl-operator keys create -k pro1


Enter password to mount: **********
Enter private key: ****************************************************************
Created key: pro1
```
3. `list` - This command will list all the keys currently present (only by file name and created date)
```
$ dcl-operator keys list

Enter password to mount: **********
   -------------------------------------------------------
   Name                           Created
   -------------------------------------------------------
   pro1                            30-04-2024 11:36:16
   -------------------------------------------------------
```
4. `delete` - This command will delete a key. This command will need a flag --key-name(-k). The name given will be searched in the `.decrypted_keys` and deleted
```
$ dcl-operator keys delete -k pro1

Enter password to mount: **********
Deleted key: pro1

$ dcl-operator keys list

Enter password to mount: **********
   -------------------------------------------------------
   Name                           Created
   -------------------------------------------------------
   -------------------------------------------------------
```

Going forward, the CLI will ask for the mount password to decrpyt and use these keys. This is how the config will look like when using encrypted keys

```
{
  "prover_private_keys": [
   "pro1",
   "pro2"
  ],
  "operator_private_key": "op1",
  "prover_registry_address": "0x91013d3CecE055603D8b1EE7DCB1f670f480fe24",
  "eth_rpc_url": "<Mainnet RPC URL>",
  "chain_id": 1,
  "gas_limit": 5000000,
  "tx_receipt_timeout": 600,
  "expiry_in_days": 1,
  "use_encrypted_keys": true
}
```

Similarly for the challenger

```
{
  "challenger_private_keys": [
   "ch1",
   "ch2"
  ],
  "operator_private_key": "op1",
  "prover_registry_address": "0xeFFE8c100029F71924554aEd382f1919ecA6b203",
  "eth_rpc_url": "<Mainnet RPC URL>",
  "chain_id": 1,
  "gas_limit": 5000000,
  "tx_receipt_timeout": 600,
  "expiry_in_days": 1,
  "use_encrypted_keys": true
}
```

