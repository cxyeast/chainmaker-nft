package myclient

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	"github.com/hokaccha/go-prettyjson"
)

func makeFile(filePath, fileName string) error {
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return fmt.Errorf("make dir failed, %s", err.Error())
	}

	f, err := os.Create(filepath.Join(filePath, fileName))
	if err != nil {
		return fmt.Errorf("create file failed, %s", err.Error())
	}

	if err = f.Close(); err != nil {
		return err
	}

	return nil
}

func printPrettyJson(data interface{}) {
	output, err := prettyjson.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
}

func CheckProposalRequestResp(resp *common.TxResponse, needContractResult bool) error {
	if resp.Code != common.TxStatusCode_SUCCESS {
		if resp.Message == "" {
			resp.Message = resp.Code.String()
		}
		return errors.New(resp.Message)
	}

	if needContractResult && resp.ContractResult == nil {
		return fmt.Errorf("contract result is nil")
	}

	if resp.ContractResult != nil && resp.ContractResult.Code != 0 {
		return errors.New(resp.ContractResult.Message)
	}

	return nil
}
