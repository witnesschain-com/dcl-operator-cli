## Storing keys in Web3 secret Storage format
[Web3 Secret Storage 
format](https://ethereum.org/en/developers/docs/data-structures-and-encoding/web3-secret-storage/) 
provides [data at rest](https://en.wikipedia.org/wiki/Data_at_rest) protection to your private keys.

## Import Keys

`init` - This command is used to initiate a new keystore file system and the type will be dependent on the key-type chosen. 

```bash
$ dcl-operator keys init
Init keystore done
```
This command creates `~/.witnesschain/cli/.w3secretkeys` directory.

`import` - Import your ECDSA keys into keystore and store them in web3 
secret storage format. The keys are stored in 
`~/.witnesschain/cli/.w3secretkeys` directory.

```bash
$ dcl-operator keys import
Enter password to import: **********
Enter private key: ****************************************************************
Imported key: 0x57a1251f30285801B2E731082259cB25DEEFA64A
```

Note: Remember the password you entered to import keys. It will be used 
later when you start watchtower.

Similarly, import other provers and challengers keys.

`list` - List all the keys stored in keystore.

```bash
$ dcl-operator keys list
   ------------------------------------------------------------------------------------------------------------------------------------------------------
   Name                                                                                                                          Created
   ------------------------------------------------------------------------------------------------------------------------------------------------------
   /home/ubuntu/.witnesschain/cli/.w3secretkeys/0x57a1251f30285801B2E731082259cB25DEEFA64A.ecdsa.key.json                  25-10-2024 15:17:25
   ------------------------------------------------------------------------------------------------------------------------------------------------------
```

Sample configuration files for registering provers and challengers:-

File: `challenger.json`

```
{
  "challenger_encrypted_keys": [
    "/home/ubuntu/.witnesschain/cli/.w3secretkeys/0x63a1251f30285801B2E731082259cB25DEEFA692.ecdsa.key.json",
    "/home/ubuntu/.witnesschain/cli/.w3secretkeys/0x57a1251f30285801B2E731082259cB25DEEFA64A.ecdsa.key.json"
  ],
  "operator_encrypted_key": "/home/ubuntu/.witnesschain/cli/.w3secretkeys/0x34a1251f30285801B2E731082259cB25DEEFA684.ecdsa.key.json ",
  "eth_rpc_url": "https://blue-orangutan-rpc.eu-north-2.gateway.fm",
}
```

File: `prover.json`

```bash
{
  "prover_encrypted_keys": [
    "/home/ubuntu/.witnesschain/cli/.w3secretkeys/0x63a1251f30285801B2E731082259cB25DEEFA692.ecdsa.key.json",
    "/home/ubuntu/.witnesschain/cli/.w3secretkeys/0x57a1251f30285801B2E731082259cB25DEEFA64A.ecdsa.key.json"
  ],
  "operator_encrypted_key": "/home/ubuntu/.witnesschain/cli/.w3secretkeys/0x34a1251f30285801B2E731082259cB25DEEFA684.ecdsa.key.json ",
  "eth_rpc_url": "https://blue-orangutan-blockscout.eu-north-2.gateway.fm/",
}
```

