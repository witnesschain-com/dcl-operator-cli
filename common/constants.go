package wc_common

const (
	DefaultOpConfig string  = "config/operator-config.json"
	EncryptedDir    string  = ".encrypted_keys"
	DecryptedDir    string  = ".decrypted_keys"
	GoCryptFSConfig string  = EncryptedDir + "/gocryptfs.conf"
	GoCryptFSDiriv  string  = EncryptedDir + "/gocryptfs.diriv"
	MinEntropyBits  float64 = 50
)
