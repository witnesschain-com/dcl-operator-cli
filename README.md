# DCL Operator CLI

## Description
dcl-operator is a command-line interface (CLI) tool for registering provers and challengers.
You can refer to the [How to use the config files](#how-to-use-the_config-files) section to understand how to use the config files.

## Installation
You can get the dcl-operator cli by building from source

Installation instructions for building from source is available in 
[docs/install.md](docs/install.md).


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
|registerWatchtower| Used to register challengers and provers |

## How to use the config files
The structure and details in the config file might differ based on the functionality you are using. Enter the following details in the config file. Changing the key names in the json file will lead to misbehavior

### For the challenger

#### related commands - registerChallenger, deRegisterChallenger
Default file: operator-challenger-config.json.template (reference file)

| Field | Description |
|----------|----------|
|challenger_private_keys | Private keys of the challengers (use this field if you want to enter raw keys) |
|challenger_encrypted_keys | Encrypted private keys of the challengers (use this field if you want to enter encrypted key names) |
|operator_private_key | Private key of the operator(on which the actions will be performed) (use this field if you want to enter raw key)|
|operator_encrypted_key | Encrypted private key of the operator(on which the actions will be performed) (use this field if you want to enter raw key)|
|eth_rpc_url | The RPC URL where you want to perform the transactions |
|gas_limit | The gas limit you want to set while sending the transactions (Default value = 1000000). No need to add in the config unless you want to overwrite the default values.  |
|tx_receipt_timeout| Timeout in seconds for waiting of tx receipts (Default value = 300). No need to add in the config unless you want to overwrite the default values. |
|expiry| Expiry in days after which the operator signature becomes invalid (Default value = 1). No need to add in the config unless you want to overwrite the default values. |

### For the prover

#### related commands - registerProver, deRegisterProver
Default file: operator-prover-config.json.template (reference file)

| Field | Description |
|----------|----------|
|prover_private_keys | Private keys of the provers (use this field if you want to enter raw keys) |
|prover_encrypted_keys | Encrypted private keys of the provers (use this field if you want to enter encrypted key names) |
|operator_private_key | Private key of the operator(on which the actions will be performed) (use this field if you want to enter raw key)|
|operator_encrypted_key | Encrypted private key of the operator(on which the actions will be performed) (use this field if you want to enter raw key)|
|eth_rpc_url | The RPC URL where you want to perform the transactions |
|encrypted_key_type | The type of encryption used for the keys (valid values = w3secretkeys/gocryptfs) (Default value = w3secretkeys). No need to add in the config unless you want to overwrite the default values.   |
|gas_limit | The gas limit you want to set while sending the transactions (Default value = 1000000). No need to add in the config unless you want to overwrite the default values.  |
|tx_receipt_timeout| Timeout in seconds for waiting of tx receipts (Default value = 300). No need to add in the config unless you want to overwrite the default values. |
|expiry| Expiry in days after which the operator signature becomes invalid (Default value = 1). No need to add in the config unless you want to overwrite the default values. |

## How to use the encrypted keys

To store and use the keys in an encrypted format, use the `challenger_encrypted_keys`, `prover_encrypted_keys` and `operator_encrypted_key` fields in the configs. If they have same values entered into it, the CLI tool will use the keys stored in the encrypted format in the given key name.

For `w3secretkeys`, the default path is `~/.witnesschain/cli/.w3secretkeys`. These are similar to EigenLayer ECDSA keys. Read more about [it](https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation#create-and-list-keys)

### Sub-commands
The `keys` command has 4 sub-commands. Let's look at them one by one -
The flags which can be used with the key commands are

| Flag | Alias | Usage |
|----------|----------|----------|
|--key-name | -k <value> | This flag is used to specify the key name |
|--key-type | -t <value> | This flag is used to specify the keystore type (w3secretkeys/gocryptfs) |
|--insecure | -i | This flag is used to bypass password strength validations |

1. `init` - This command is used to initiate a new keystore file system and the type will be dependent on the key-type chosen. 

```
init operation for key-type gocryptfs
$ dcl-operator keys init -t w3secretkeys -i

Creating directory:  .w3secretkeys
Init keystore done
```
Contrary to `gocryptfs` type, `w3secretkeys` does not need a password during init operation. Instead you will be asked for password while you create/import/export/delete keys. The key difference between these two types is `gocryptfs` uses the password used during the init operation. And with `w3secretkeys` type, individual keys can have their own password. You will need that password while you create/import/export keys. For consistency when using multiple keys in the config, it is mandatory for all the `w3secretkeys` keys to have the same password.

2. `create` - This command will create a new key. The value you pass with the `-k` flag will be the name of the key which will be referred in the future by the CLI. This name should be mentioned in the config file to extract the corresponding private key. When you run this command, it will ask for password to mount and then ask you to enter the private key
```
4. `list` - This command will list all the keys currently present (only by file name and created date)
```
$ dcl-operator keys list -t w3secretkeys
   -------------------------------------------------------
   Name                           Created
   -------------------------------------------------------
   pro1                            30-04-2024 11:36:16
   -------------------------------------------------------
```



The below example shows how you can use the key names which will be taken from the default path

```
{
  "prover_private_keys": [
    "<raw-prover-private-key>"
  ],
  "operator_private_key": "<raw-operator-private-key>",
  "eth_rpc_url": "<Mainnet RPC URL>",
}
```

Similarly for the challenger

```
{
  "challenger_private_keys": [
    "<raw-challenger-private-key>"
  ],
  "operator_private_key": "<raw-operator-private-key>",
  "eth_rpc_url": "<Mainnet RPC URL>",
}
```

The below example shows how you can use encrypted keys using `w3secretkeys` type. The type can also be `gocryptfs`
```
{
  "prover_encrypted_keys": [
    "~/alternate/path/to/your/keys/pro1"
  ],
  "operator_encrypted_key": "~/alternate/path/to/your/keys/op1",
  "encrypted_key_type": "w3secretkeys",
  "eth_rpc_url": "<Mainnet RPC URL>"
}
```
Similarly for the challenger

```
{
  "challenger_encrypted_keys": [
    "~/alternate/path/to/your/keys/ch1"
  ],
  "operator_encrypted_key": "~/alternate/path/to/your/keys/op1",
  "encrypted_key_type": "w3secretkeys",
  "eth_rpc_url": "<Mainnet RPC URL>",
}
```

The below example shows how you can use the key names which will be taken from an alternate path using `gocryptfs` type. The type can also be `w3secretkeys`
```
{
  "prover_encrypted_keys": [
    "~/alternate/path/to/your/keys/.encrypted_keys/pro1"
  ],
  "operator_encrypted_key": "~/alternate/path/to/your/keys/.encrypted_keys/op1",
  "encrypted_key_type": "gocryptfs",
  "eth_rpc_url": "<Mainnet RPC URL>"
}
```
Similarly for the challenger

```
{
  "challenger_encrypted_keys": [
    "~/alternate/path/to/your/keys/.encrypted_keys/ch1"
  ],
  "operator_encrypted_key": "~/alternate/path/to/your/keys/.encrypted_keys/op1",
  "encrypted_key_type": "gocryptfs",
  "eth_rpc_url": "<Mainnet RPC URL>"
}
```
