package sec

import (
	"fmt"
	"io"
	"os"

	enc "github.com/named-data/ndnd/std/encoding"
	"github.com/named-data/ndnd/std/ndn"
	"github.com/named-data/ndnd/std/object"
	"github.com/named-data/ndnd/std/security"
	"github.com/named-data/ndnd/std/security/keychain"
	sig "github.com/named-data/ndnd/std/security/signer"
	"github.com/spf13/cobra"
)

func keychainList(_ *cobra.Command, args []string) {
	kc, err := keychain.NewKeyChain(args[0], object.NewMemoryStore())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open keychain: %s\n", err)
		os.Exit(1)
		return
	}

	for _, id := range kc.Identities() {
		fmt.Printf("%s\n", id.Name())
		for _, key := range id.Keys() {
			fmt.Printf("==> %s\n", key.KeyName())
		}
	}
}

func keychainImport(_ *cobra.Command, args []string) {
	kc, err := keychain.NewKeyChain(args[0], object.NewMemoryStore())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open keychain: %s\n", err)
		os.Exit(1)
		return
	}

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read input: %s\n", err)
		os.Exit(1)
		return
	}

	err = keychain.InsertFile(kc, input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert keychain entries: %s\n", err)
		os.Exit(1)
		return
	}
}

func keychainExport(_ *cobra.Command, args []string) {
	name, err := enc.NameFromStr(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid key name: %s\n", args[1])
		os.Exit(1)
		return
	}

	kc, err := keychain.NewKeyChain(args[0], object.NewMemoryStore())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open keychain: %s\n", err)
		os.Exit(1)
		return
	}

	keyName := name
	id, err := security.GetIdentityFromKeyName(name)
	if err != nil { // not a key name
		id = name
		keyName = nil
	}

	idObj := kc.IdentityByName(id)
	if idObj == nil {
		fmt.Fprintf(os.Stderr, "Identity not found: %s\n", id)
		os.Exit(1)
		return
	}

	var signer ndn.Signer
	if keyName == nil {
		if len(idObj.Keys()) > 0 {
			signer = idObj.Keys()[0].Signer()
		}
	} else {
		for _, key := range idObj.Keys() {
			if key.KeyName().Equal(keyName) {
				signer = key.Signer()
				break
			}
		}
	}
	if signer == nil {
		fmt.Fprintf(os.Stderr, "Key not found: %s\n", keyName)
		os.Exit(1)
		return
	}

	secret, err := sig.MarshalSecret(signer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to encode secret key: %s\n", err)
		os.Exit(1)
		return
	}

	out, err := security.PemEncode(secret.Join())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to convert secret key to text: %s\n", err)
		os.Exit(1)
		return
	}

	os.Stdout.Write(out)
}
