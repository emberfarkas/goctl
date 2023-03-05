package eth

import (
	"context"
	"fmt"
	"math/big"

	"github.com/emberfarkas/goctl/internal/eth/adorn"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-bamboo/pkg/log"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

// 这个工具主要用来测试eth相关的借口

// Cmd represents the config command
var (
	adornCmd = &cobra.Command{
		Use:   "adorn",
		Short: "adorn相关",
		Long:  `获取adorn详情`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return getAdorn(cmd.Context())
		},
	}
	to int64
)

func init() {

	// Here you will define your flags and configuration settings.

	adornCmd.Flags().Int64VarP(&to, "to", "t", 26171426, "到哪里")
}

var tokens map[string]*big.Int = map[string]*big.Int{}
var accounts map[string]map[string]*big.Int = map[string]map[string]*big.Int{}

func getAdorn(ctx context.Context) error {
	//bsc := "https://bsc-dataseed4.ninicoin.io"
	//eth := "https://ethereum.blockpi.network/v1/rpc/public"
	//bsc := "https://falling-blue-borough.bsc-testnet.quiknode.pro/5e10cf573ba10d2c1bf2fe0bf9dd3fdbee8218d6/"
	bsc := "https://attentive-necessary-sound.bsc.quiknode.pro/c2bdc463b449995c2e8dc5bb359f869159858b85/"
	rpc, err := ethclient.Dial(bsc)
	if err != nil {
		return err
	}
	log.Infof("-------------------")
	inst, err := adorn.NewErc1155(common.HexToAddress("0x2DFEb752222ccceCB9BC0a934b02C3A86f633900"), rpc)

	// 26171426
	step := int64(10000)
	for i := int64(22388143); i < 26171426; {
		log.Debugf("from: %v, to: %v", i, i+step)
		end := uint64(i + step)
		iter1, err := inst.FilterTransferSingle(&bind.FilterOpts{
			Start: uint64(i),
			End:   &end,
		}, nil, nil, nil)
		if err != nil {
			return err
		}
		for iter1.Next() {
			if err := saveAccountSingle(iter1.Event); err != nil {
				return err
			}
		}

		iter2, err := inst.FilterTransferBatch(&bind.FilterOpts{
			Start: uint64(i),
			End:   &end,
		}, nil, nil, nil)
		if err != nil {
			return err
		}
		for iter2.Next() {
			if err := saveAccountBatch(iter2.Event); err != nil {
				return err
			}
		}
		i = i + step
	}
	return saveXlsx()
}

func saveAccountSingle(single *adorn.Erc1155TransferSingle) error {
	if err := sunFrom(single.From, single.Id, single.Value); err != nil {
		return err
	}
	if err := addTo(single.To, single.Id, single.Value); err != nil {
		return err
	}
	return nil
}

func saveAccountBatch(batch *adorn.Erc1155TransferBatch) error {
	for i := 0; i < len(batch.Ids); i++ {
		id := batch.Ids[i]
		val := batch.Values[i]
		if err := sunFrom(batch.From, id, val); err != nil {
			return err
		}
		if err := addTo(batch.To, id, val); err != nil {
			return err
		}
	}
	return nil
}

func sunFrom(from common.Address, id *big.Int, value *big.Int) error {
	var d, ok = accounts[from.Hex()]
	if !ok {
		d = map[string]*big.Int{}
		accounts[from.Hex()] = d
	}
	val1, ok := d[id.String()]
	if !ok {
		val1 = big.NewInt(0)
	}
	val1 = new(big.Int).Sub(val1, value)
	d[id.String()] = val1
	return nil
}

func addTo(to common.Address, id *big.Int, value *big.Int) error {
	var d, ok = accounts[to.Hex()]
	if !ok {
		d = map[string]*big.Int{}
		accounts[to.Hex()] = d
	}
	var val1, ok1 = d[id.String()]
	if !ok1 {
		val1 = big.NewInt(0)
	}
	val1 = big.NewInt(0).Add(val1, value)
	d[id.String()] = val1
	return nil
}

func saveXlsx() error {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A1", "地址")
	f.SetCellValue("Sheet2", "B1", "数量")

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	offset := 2
	for key, val := range accounts {
		f.SetCellValue("Sheet2", fmt.Sprintf("A%d", offset), key)
		cnt := int64(0)
		for _, b := range val {
			cnt = b.Int64() + cnt
		}
		f.SetCellValue("Sheet2", fmt.Sprintf("B%d", offset), cnt)
		offset++
	}
	return f.SaveAs("./Book1.xlsx")
}
