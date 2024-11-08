package operator_commands

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/publicsuffix"
)


const broker_url = "https://api.witnesschain.com"

type PolClaim struct {
	Country string `json:"country"`
	City string `json:"city"`
	Region string `json:"region"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Radius float32 `json:"radius"`
}

type PublicKey struct{
	Solana string `json:"solana"`
	Ethereum string `json:"ethereum"`
}

type PreLoginBody struct{
	Publickey string `json:"publicKey"`
	KeyType string `json:"keyType"`
	Role string `json:"role"`
	ProjectName string `json:"projectName"`
	Claims PolClaim `json:"claims"`
	WalletPublicKey PublicKey `json:"walletPublicKey"`
	ClientVersion string `json:"clientVersion"`
}

type PreLoginMessage struct {
	Message string `json:"message"`
}

type PreloginResult struct {
	Result PreLoginMessage `json:"result"`
}

type LoginBody struct {
	Signature string `json:"signature"`
}

type LoginResponse struct {
	Result LoginResult `json:"result"`
}

type LoginResult struct {
	Success bool `json:"success"`
}

type ChallengerParameters struct {
	Id string `json:"id"`
}

type ChallengerResponse struct {
	Result ChallengerResult `json:"result"`
}

type ChallengerResult struct {
	Id string `json:"id"`
	LastAlive string `json:"last_alive"`
}

type ProverResponse struct {
	Result ProverResult `json:"result"`
}

type ProverResult struct {
	Id string `json:"id"`
	LastAlive string `json:"last_alive"`
}

func WatchtowerStatusCmd() *cli.Command {
	return &cli.Command{
		Name:  "watchtowerStatus",
		Usage: "Find status of watchtower",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "watchtower-address",
				Required: true,
			},
			&cli.BoolFlag{
				Name: "pob",

			},
			&cli.BoolFlag{
				Name: "pol",
			},

		},
		Action: func(cCtx *cli.Context) error {
			WatchtowerStatus(cCtx.String("watchtower-address"))
			return nil
		},
	}
}

func WatchtowerStatus(watchtowerAddress string) {
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("Watchtower:\t\t%s\n", watchtowerAddress)
	fmt.Println(strings.Repeat("=", 80))
	
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
			log.Fatal(err)
		}

	client := http.Client{
		Jar: jar,
	}

	challenger(client, "IPv4",  watchtowerAddress)
	challenger(client, "IPv6",  watchtowerAddress)
	watchtowerType := WatchtowerType(client, watchtowerAddress)
	fmt.Printf("Watchtower Type:\t%s\n", watchtowerType)
	fmt.Println(strings.Repeat("=", 80))
}

func challenger(client http.Client, protocol string, id string) bool{
	challengerParameters := ChallengerParameters{
		Id: protocol + "/" + id,
	}

	challengerParametersBytes, err := json.Marshal(challengerParameters)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Post("https://api.witnesschain.com/proof/v1/pol/challenger", "application/json", bytes.NewBuffer(challengerParametersBytes))

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	challengerResponse := ChallengerResponse{}
	json.Unmarshal(body, &challengerResponse)

	if len(challengerResponse.Result.LastAlive) == 0 {
		fmt.Printf("%s:\t\t\tNot found\n", protocol)
		return false
	}

	last_alive, err := time.Parse(time.RFC3339, challengerResponse.Result.LastAlive)
	if err != nil {
		log.Fatal(err)
	}

	last_alive_duration := time.Now().Sub(last_alive).Round(time.Second)


	if (last_alive_duration > time.Duration(time.Minute*5)){
		color.Red("%s:\t\t\tInactive:\t\tlast seen %s ago", protocol, last_alive_duration)
		return false
	}

	color.Green("%s:\t\t\tActive:\tlast seen %s ago", protocol, last_alive_duration)
	return true
}

func WatchtowerType(client http.Client,watchtowerAddress string) string{
	challengerParameters := ChallengerParameters{
		Id: "IPv4/" + watchtowerAddress,
	}

	challengerParametersBytes, err := json.Marshal(challengerParameters)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Post("https://api.witnesschain.com/proof/v1/pol/challenger", "application/json", bytes.NewBuffer(challengerParametersBytes))

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	challengerResponse := ChallengerResponse{}
	json.Unmarshal(body, &challengerResponse)

	if len(challengerResponse.Result.LastAlive) != 0 {
		return "Challenger"
	}

	proverParameters := ChallengerParameters{
		Id: "IPv4/" + watchtowerAddress,
	}

	proverParametersBytes, err := json.Marshal(proverParameters)
	if err != nil {
		log.Fatal(err)
	}

	response, err = client.Post("https://api.witnesschain.com/proof/v1/pol/prover", "application/json", bytes.NewBuffer(proverParametersBytes))

	body, err = io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	proverRespone := ProverResponse{}
	json.Unmarshal(body, &proverRespone)
	if len(proverRespone.Result.LastAlive) != 0 {
		return "Prover"
	}

	return "Unable to determine watchtower status"

}

func prelogin(client http.Client, privateKey *ecdsa.PrivateKey) string {

	loginAddress := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	log.SetFlags(log.LstdFlags | log.Lshortfile) 
	pre_login_body := PreLoginBody{
		KeyType: "ethereum",
		Role: "prover",
		ProjectName: "witness",
		ClientVersion: "99999999999",
		Publickey: loginAddress,
		WalletPublicKey: PublicKey{
			Ethereum: loginAddress,
		},
		Claims: PolClaim{
			Latitude: 0.0,
			Longitude: 0.0,
			Radius: 1000000.0,
			Country: "IN",
			City: "Bengaluru",
			Region: "Karnataka",
		},
	}
	
	body, err:= json.Marshal(pre_login_body)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Post("https://api.witnesschain.com/proof/v1/pol/pre-login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	result := PreloginResult{}

	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result.Result.Message
}


func login(client http.Client, privateKey *ecdsa.PrivateKey, message string) bool{
	signature := signMessage(privateKey, message)

	loginBody := LoginBody{
		Signature: signature,
	}

	loginBodyBytes, err := json.Marshal(loginBody)
	
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Post("https://api.witnesschain.com/proof/v1/pol/login", "application/json", bytes.NewBuffer(loginBodyBytes))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	loginResponse := LoginResponse{}
	json.Unmarshal(body, &loginResponse)
	if loginResponse.Result.Success != true {
		log.Fatal("unable to login to witnesschain")
	}
	return loginResponse.Result.Success
}


func signMessage(privateKey *ecdsa.PrivateKey, message string) (string) {

	prefixed_msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(prefixed_msg))
	sig, _ := crypto.Sign(hash.Bytes(), privateKey)
	sig[64] += 27
	result := "0x" + hex.EncodeToString(sig)
	return result
}
