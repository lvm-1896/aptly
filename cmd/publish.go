package cmd

import (
	"github.com/aptly-dev/aptly/pgp"
	"github.com/smira/commander"
	"github.com/smira/flag"
)

func getSigner(flags *flag.FlagSet) (pgp.Signer, error) {
	signer := context.GetSigner()
	conf := context.Config()
	if LookupOption(conf.GpgDisableSign, flags, "skip-signing") {
		return nil, nil
	}
	signer.SetKey(LookupOptionString(conf.GpgSigningKey, flags, "gpg-key"))
	signer.SetKey(LookupOptionString(conf.GpgSigningKey, flags, "gpg-key"))
	signer.SetKeyRing(LookupOptionString(conf.GpgKeyring, flags, "keyring"),
		LookupOptionString(conf.GpgSecretKeyring, flags, "secret-keyring"))
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
