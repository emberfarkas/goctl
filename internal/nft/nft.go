package nft

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-bamboo/contrib/contracts/erc721"
	"github.com/go-bamboo/pkg/fs/s3"
	"github.com/go-bamboo/pkg/log"
	"github.com/imroc/req/v3"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var Cmd = &cobra.Command{
	Use:   "nft",
	Short: "nft相关辅助工具",
	Long:  `一些批处理nft的工具`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run1(cmd.Context())
	},
}

var contract string
var tokenID string

func init() {
	Cmd.Flags().StringVar(&contract, "contract", "", "")
	Cmd.Flags().StringVar(&tokenID, "tokenID", "", "")
}

func run1(ctx context.Context) error {
	contractAddress := common.HexToAddress(contract)

	rpc := os.Getenv("RPC")

	gClient, err := ethclient.Dial(rpc)
	if err != nil {
		return err
	}
	client, err := erc721.NewErc721(contractAddress, gClient)
	if err != nil {
		return err
	}
	uri, err := client.BaseUri(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		log.Errorw("metadata", "contract", contractAddress.Hex(), "err", err)
		return err
	}
	baseURI, err := url.Parse(strings.TrimSpace(uri))
	if err != nil {
		return err
	}
	tokenURI := baseURI.JoinPath(fmt.Sprintf("%v.json", tokenID))
	log.Infof("url: %v", tokenURI)
	res, err := req.R().EnableDump().Get(tokenURI.String())
	if err != nil {
		return err
	}
	md := res.Bytes()
	imageURL := gjson.GetBytes(md, "image").String()
	mainURL := gjson.GetBytes(md, "modified_url").String()

	log.Infof("imageUrl: %v", imageURL)
	log.Infof("mainURL: %v", mainURL)

	return updateToken(ctx, md)
}

func updateToken(ctx context.Context, md []byte) error {
	// 下面更新数据
	newImageUrl := fmt.Sprintf("https://ipfs-v2.halonft.art/bsc_v2/haloworld/L1/image/%v.jpg", tokenID)
	nmd, err := sjson.SetBytes(md, "image", newImageUrl)
	if err != nil {
		return err
	}
	newMainUrl := fmt.Sprintf("https://ipfs-v2.halonft.art/bsc_v2/haloworld/L1/res/%v.main", tokenID)
	nmd, err = sjson.SetBytes(nmd, "modified_url", newMainUrl)
	if err != nil {
		return err
	}
	log.Debugf("%v", string(nmd))

	s3session, err := s3.New(&s3.Conf{})
	if err != nil {
		log.Error(err)
		return err
	}
	jsonURI, err := s3session.UploadBytesToBucketDir(ctx, "ipfs-v2.halonft.art", "bsc_v2/haloworld/L1/token", fmt.Sprintf("%v.json", tokenID), nmd)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debugf(jsonURI)
	return nil
}
