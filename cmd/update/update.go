package update

import (
	"log"

	"github.com/gvcgo/wgcf/cloudflare"
	. "github.com/gvcgo/wgcf/cmd/shared"
	"github.com/gvcgo/wgcf/config"
	"github.com/gvcgo/wgcf/util"
	"github.com/gvcgo/wgcf/wireguard"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deviceName string

var shortMsg = "Updates the current Cloudflare Warp account, preparing it for connection"

var Cmd = &cobra.Command{
	Use:   "update",
	Short: shortMsg,
	Long: FormatMessage(shortMsg, `
If a new/different license key is provided, the current device will be bound to the new key and its parent account. 
Please note that there is a maximum limit of 5 active devices linked to the same account at a given time.`),
	Run: func(cmd *cobra.Command, args []string) {
		if err := updateAccount(); err != nil {
			log.Fatal(util.GetErrorMessage(err))
		}
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&deviceName, "name", "n", "", "Device name displayed under the 1.1.1.1 app")
}

func updateAccount() error {
	if !IsConfigValidAccount() {
		return errors.New("no account detected")
	}

	ctx := CreateContext()
	thisDevice, err := cloudflare.GetSourceDevice(ctx)
	if err != nil {
		return err
	}
	_, thisDevice, err = ensureLicenseKeyUpToDate(ctx, thisDevice)
	if err != nil {
		return err
	}

	boundDevice, err := cloudflare.GetSourceBoundDevice(ctx)
	if err != nil {
		return err
	}
	if boundDevice.Name == nil || (deviceName != "" && deviceName != *boundDevice.Name) {
		log.Println("Setting device name")
		if _, err := SetDeviceName(ctx, deviceName); err != nil {
			return err
		}
	}

	boundDevice, err = cloudflare.UpdateSourceBoundDeviceActive(ctx, true)
	if err != nil {
		return err
	}
	if !boundDevice.Active {
		return errors.New("failed activating device")
	}

	PrintDeviceData(thisDevice, boundDevice)
	log.Println("Successfully updated Cloudflare Warp account")
	return nil
}

func ensureLicenseKeyUpToDate(ctx *config.Context, thisDevice *cloudflare.Device) (*cloudflare.Account, *cloudflare.Device, error) {
	if thisDevice.Account.License != ctx.LicenseKey {
		log.Println("Updated license key detected, re-binding device to new account")
		return updateLicenseKey(ctx)
	}
	return nil, thisDevice, nil
}

func updateLicenseKey(ctx *config.Context) (*cloudflare.Account, *cloudflare.Device, error) {
	newPrivateKey, err := wireguard.NewPrivateKey()
	if err != nil {
		return nil, nil, err
	}
	newPublicKey := newPrivateKey.Public()
	if _, _, err := cloudflare.UpdateLicenseKey(ctx, newPublicKey.String()); err != nil {
		return nil, nil, err
	}

	viper.Set(config.PrivateKey, newPrivateKey.String())
	if err := viper.WriteConfig(); err != nil {
		return nil, nil, err
	}

	account, err := cloudflare.GetAccount(ctx)
	if err != nil {
		return nil, nil, err
	}
	thisDevice, err := cloudflare.GetSourceDevice(ctx)
	if err != nil {
		return nil, nil, err
	}

	if account.License != ctx.LicenseKey {
		return nil, nil, errors.New("failed to update license key")
	}
	if thisDevice.Key != newPublicKey.String() {
		return nil, nil, errors.New("failed to update public key")
	}

	return account, thisDevice, nil
}
