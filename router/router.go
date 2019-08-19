package router

import (
	"fmt"

	"github.com/ditcraft/cli/api"
	"github.com/ditcraft/cli/config"
	"github.com/ditcraft/cli/ethereum"
	"github.com/ditcraft/cli/helpers"
)

// GetBalances will print out the xDAI/xDIT and KNW-token balances
func GetBalances() error {
	var nativeToken float64
	var labelArray []string
	var balanceArray []float64
	var err error

	nativeToken, labelArray, balanceArray, err = api.GetBalances()
	if err != nil {
		nativeToken, labelArray, balanceArray, err = ethereum.GetBalances()
		if err != nil {
			return err
		}
	}

	helpers.PrintLine(fmt.Sprintf(config.DitConfig.Currency+"-Balance: %.2f "+config.DitConfig.Currency, nativeToken), helpers.INFO)
	helpers.PrintLine("", helpers.INFO)

	helpers.PrintLine("KNW-Balances:", helpers.INFO)
	atLeastOneBalance := false
	for i := 0; i < len(labelArray); i++ {
		if balanceArray[i] >= 0.01 {
			atLeastOneBalance = true
			helpers.PrintLine(fmt.Sprintf(" - "+labelArray[i]+": %.2f KNW", balanceArray[i]), helpers.INFO)
		}
	}

	if !atLeastOneBalance {
		helpers.PrintLine("- No labels with any balance found", helpers.INFO)
	}
	helpers.PrintLine("", helpers.INFO)
	helpers.PrintLine("URL: https://explorer.ditcraft.io/address/"+config.DitConfig.EthereumKeys.Address, helpers.INFO)

	return nil
}
