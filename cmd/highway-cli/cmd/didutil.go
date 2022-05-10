package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/spf13/cobra"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name: "goal",
		Prompt: &survey.Select{
			Message: "Choose an option:",
			Options: []string{"Create a new DID", "Update existing DID"},
			Default: "Create a new DID",
		},
	},
	{
		Name:      "name",
		Prompt:    &survey.Input{Message: "What is your name?"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},

	{
		Name:   "age",
		Prompt: &survey.Input{Message: "How old are you?"},
	},
}

// DidUtilCmd returns did cobra Command.
func DidUtilCmd() *cobra.Command {
	// the answers will be written to this struct
	answers := struct {
		Goal          string // survey will match the question and field names
		FavoriteColor string `survey:"color"` // or you can tag fields to match a specific name
		Age           int    // if the types don't match, survey will convert it
	}{}

	// run the survey
	cmd := &cobra.Command{
		Use:   "did",
		Short: "Generate DID Documents and handle WebAuthn",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// run the survey
			selectOpt := &survey.Select{
				Message: "Choose an option:",
				Options: []string{"Create a new DID", "Update existing DID"},
				Default: "Create a new DID",
			}
			// perform the questions
			survey.AskOne(selectOpt, &answers.Goal, nil)

			if answers.Goal == "Create a new DID" {
				err = surveyNewDid()
			} else {
				err = surveyExistingDid()
			}

			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func surveyNewDid() error {
	// Webauthn confirmation
	webauthn := false
	prompt := &survey.Confirm{
		Message: "Would you like to generate a Webauthn Credential in your browser?",
	}
	survey.AskOne(prompt, &webauthn)

	if webauthn {
		// Open browser and serve index.html
	} else {
		pubKey := ""
		// Present public key types
		selectOpt := &survey.Select{
			Message: "What type of key would you like to generate",
			Options: []string{
				"Ed25519VerificationKey2018",
				"EcdsaSecp256k1VerificationKey2019",
				"RsaVerificationKey2018",
			},
			Default: "Ed25519VerificationKey2018",
		}
		// perform the questions
		survey.AskOne(selectOpt, &pubKey, nil)

		// Parse Account Addr
		addr := genAccAddressDid()
		didRoot, err := did.ParseDID(addr)
		if err != nil {
			return err
		}

		var vm *did.VerificationMethod

		// Parse DID Controller
		didController, err := did.ParseDID(fmt.Sprintf("%s#test", didRoot))
		if err != nil {
			return err
		}

		// Switch on pubKey
		switch pubKey {
		case "EcdsaSecp256k1VerificationKey2019":
			// Generate EcdsaSecp256k1VerificationKey2019
			priv := secp256k1.GenPrivKey()
			pub := priv.PubKey()
			vm, _ = did.NewVerificationMethod(*didRoot, ssi.ECDSASECP256K1VerificationKey2019, *didController, pub)
			// did.ParseDID(input string)
		case "RsaVerificationKey2018":
			// Generate RsaVerificationKey2018
			priv, err := rsa.GenerateKey(rand.Reader, 2048)
			if err != nil {
				return err
			}
			pub := priv.PublicKey
			vm, _ = did.NewVerificationMethod(*didRoot, ssi.RSAVerificationKey2018, *didController, pub)
			// did.ParseDID(input string)
		default:
			pub := ed25519.GenPrivKey().PubKey()
			vm, _ = did.NewVerificationMethod(*didRoot, ssi.ED25519VerificationKey2018, *didController, pub)
			// did.ParseDID(input string)
		}

		ctx, err := ssi.ParseURI("https://www.w3.org/ns/did/v1")
		if err != nil {
			return err
		}

		// Document properties
		doc := &did.Document{
			Context: []ssi.URI{*ctx},
			ID:      *didRoot,
			Controller: []did.DID{
				*didController,
			},
		}
		doc.AddAuthenticationMethod(vm)
		buf, err := doc.MarshalJSON()
		if err != nil {
			return err
		}

		saveFile := false
		prompt := &survey.Confirm{
			Message: "Would you like to save the JSON output to disk?",
		}
		survey.AskOne(prompt, &saveFile)

		if saveFile {
			// Save to file
			dir, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			path := filepath.Join(dir, "Desktop", "did.json")
			ioutil.WriteFile(path, buf, 0644)
			fmt.Printf("Saved DID JSON Document to: %s\n", path)
		} else {
			fmt.Println(string(buf))
		}
	}
	return nil
}

func surveyExistingDid() error {
	file := ""
	prompt := &survey.Input{
		Message: "Enter the path to your DID Json file:",
		Suggest: func(toComplete string) []string {
			// Save to file
			dir, err := os.UserHomeDir()
			if err != nil {
				return nil
			}

			path := filepath.Join(dir, "Desktop")
			files, _ := filepath.Glob(path + toComplete + "*")
			return files
		},
	}
	survey.AskOne(prompt, &file)

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	doc := &did.Document{}
	err = doc.UnmarshalJSON(buf)
	if err != nil {
		return err
	}

	opt := ""
	// Present public key types
	selectOpt := &survey.Select{
		Message: "What action would you like to perform?",
		Options: []string{
			"Add Assertion Method",
			"Add Invocation Method",
			"Add Verification Method",
			"Encrypt JWE",
			"Decrypt JWE",
		},
		Default: "Add Assertion Method",
	}

	// perform the questions
	survey.AskOne(selectOpt, &opt, nil)

	// Handle opt
	switch opt {
	case "Add Assertion Method":
		// Add Assertion Method
		controller := fmt.Sprintf("%s#test-%d", doc.ID, len(doc.Controller))
		didController, err := did.ParseDID(controller)
		if err != nil {
			return err
		}
		priv := secp256k1.GenPrivKey()
		pub := priv.PubKey()
		vm, _ := did.NewVerificationMethod(doc.ID, ssi.ECDSASECP256K1VerificationKey2019, *didController, pub)
		doc.AddAssertionMethod(vm)
	case "Add Invocation Method":
		// Add Invocation Method
		controller := fmt.Sprintf("%s#test-%d", doc.ID, len(doc.Controller))
		didController, err := did.ParseDID(controller)
		if err != nil {
			return err
		}
		priv := secp256k1.GenPrivKey()
		pub := priv.PubKey()
		vm, _ := did.NewVerificationMethod(doc.ID, ssi.ECDSASECP256K1VerificationKey2019, *didController, pub)
		doc.AddCapabilityInvocation(vm)
	case "Add Verification Method":
		// Add Verification Method
		controller := fmt.Sprintf("%s#test-%d", doc.ID, len(doc.Controller))
		didController, err := did.ParseDID(controller)
		if err != nil {
			return err
		}
		priv := secp256k1.GenPrivKey()
		pub := priv.PubKey()
		vm, _ := did.NewVerificationMethod(doc.ID, ssi.ECDSASECP256K1VerificationKey2019, *didController, pub)
		doc.AddAuthenticationMethod(vm)
	case "Encrypt JWE":
		// Generate JWE
		genIpld := false
		prompt := &survey.Confirm{
			Message: "Would you like to create Object from IPLD schema?",
		}
		survey.AskOne(prompt, &genIpld)
	case "Decrypt JWE":
		text := ""
		prompt := &survey.Multiline{
			Message: "Paste your JWE object here:",
		}
		survey.AskOne(prompt, &text)
		// Decrypt JWE
		buf, err := doc.DecryptJWE(text)
		if err != nil {
			return err
		}
		fmt.Println(string(buf))
	default:
		return errors.New("Invalid option")
	}

	return nil
}

// AccAddress returns a sample account address
func genAccAddressDid() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	addrStr := sdk.AccAddress(addr).String()
	addrStr = strings.TrimLeft(addrStr, "snr")
	addrStr = strings.TrimRight(addrStr, "cosmos")
	return fmt.Sprintf("did:snr:%s", addrStr)
}
