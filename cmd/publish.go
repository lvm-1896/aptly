package cmd

import (
	"github.com/aptly-dev/aptly/pgp"
	"github.com/smira/commander"
	"github.com/smira/flag"
)

func getSigner(flags *flag.FlagSet) (pgp.Signer, error) {
	if LookupOption(context.Config().GpgDisableSign, flags, "skip-signing") {
		return nil, nil
	}

	signer := context.GetSigner()
	conf := context.Config()
	var key, flag string
	key = conf.GpgSigningKey
	flag = flags.Lookup("gpg-key").Value.String()
	if len(flag) > 0 {
		key = flag
	}
	signer.SetKey(key)
	key = conf.GpgKeyring
	flag = flags.Lookup("keyring").Value.String()
	if len(flag) > 0 {
		key = flag
	}
	kr := key
	key = conf.GpgSecretKeyring
	flag = flags.Lookup("secret-keyring").Value.String()
	if len(flag) > 0 {
		key = flag
	}
	skr := key
	signer.SetKeyRing(kr, skr)
	signer.SetPassphrase(flags.Lookup("passphrase").Value.String(), flags.Lookup("passphrase-file").Value.String())
	signer.SetBatch(flags.Lookup("batch").Value.Get().(bool))

	err := signer.Init()
	if err != nil {
		return nil, err
	}

	return signer, nil

}

func makeCmdPublish() *commander.Command {
	return &commander.Command{
		UsageLine: "publish",
		Short:     "manage published repositories",
		Subcommands: []*commander.Command{
			makeCmdPublishDrop(),
			makeCmdPublishList(),
			makeCmdPublishRepo(),
			makeCmdPublishSnapshot(),
			makeCmdPublishSwitch(),
			makeCmdPublishUpdate(),
			makeCmdPublishShow(),
		},
	}
}
