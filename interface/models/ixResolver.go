package models

import (
	"encoding/json"
	"fmt"
	"interface/config"
	"strconv"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

/*
InitLedger
InspectRequest - id int64, basicInfo BasicInfo
InspectResult - inspectionID string, detailInfo DetailInfo, images Images, etc string
QueryInspectionResult - inspectionID string
QueryAllInspectionRequest
*/
func InspectionInitLedger(pc config.PeerConfig) {
	_, commit, err := pc.InspectionContract.SubmitAsync("InitLedger")
	status, err := commit.Status()
	if err != nil {
		panic(fmt.Errorf("failed to get Inspection InitLedger transaction commit status: %w", err))
	}

	if !status.Successful {
		panic(fmt.Errorf("failed to commit Inspection InitLedger transaction with status code %v", status.Code))
	}

	fmt.Println("\n*** Inspection InitLedger committed successfully")
}
func InspectRequest(basicInfo BasicInfo, pc config.PeerConfig) Inspection {
	basicInfoJSON, err := json.Marshal(basicInfo)
	result, commit, err := pc.InspectionContract.SubmitAsync("InspectRequest",
		client.WithArguments(string(basicInfoJSON)))
	if err != nil {
		panic(fmt.Errorf("failed to submit InspectRequest transaction: %w", err))
	}

	status, err := commit.Status()
	if err != nil {
		panic(fmt.Errorf("failed to get InspectRequest transaction commit status: %w", err))
	}

	if !status.Successful {
		panic(fmt.Errorf("failed to commit InspectRequest transaction with status code %v", status.Code))
	}

	fmt.Println("\n*** InspectRequest committed successfully")
	fmt.Printf("resultJSON : %s \n", result)

	var resultStruct Inspection // YourResultStruct는 결과를 언마샬링할 구조체로 대체
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", resultStruct, err))
	}
	fmt.Printf("resultStruct %+v", resultStruct)
	return resultStruct
}

func InspectResult(id int64, detailInfo DetailInfo, images Images, etc string, pc config.PeerConfig) Inspection {
	detailInfoJSON, err := json.Marshal(detailInfo)
	imagesJSON, err := json.Marshal(images)
	result, commit, err := pc.InspectionContract.SubmitAsync("InspectResult",
		client.WithArguments(strconv.FormatInt(id, 10), string(detailInfoJSON), string(imagesJSON), etc))
	if err != nil {
		panic(fmt.Errorf("failed to submit InspectResult transaction: %w", err))
	}

	status, err := commit.Status()
	if err != nil {
		panic(fmt.Errorf("failed to get InspectResult transaction commit status: %w", err))
	}

	if !status.Successful {
		panic(fmt.Errorf("failed to commit InspectResult transaction with status code %v", status.Code))
	}

	fmt.Println("\n*** InspectResult committed successfully")
	fmt.Printf("resultJSON : %s \n", result)

	var resultStruct Inspection // YourResultStruct는 결과를 언마샬링할 구조체로 대체
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", resultStruct, err))
	}
	fmt.Printf("resultStruct %+v", resultStruct)
	return resultStruct
}

func QueryInspectResult(id string, pc config.PeerConfig) Inspection {
	result, err := pc.InspectionContract.EvaluateTransaction("QueryInspectionResult", id)
	if err != nil {
		panic(fmt.Errorf("failed to query QueryInspectionResult: %w", err))
	}

	fmt.Println("\n*** QueryInspectResult successful")
	fmt.Printf("resultJSON : %s \n", result)

	var resultStruct Inspection
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", result, err))
	}

	fmt.Printf("resultStruct %+v", resultStruct)
	return resultStruct
}

func QueryAllInspections(pc config.PeerConfig) []Inspection {
	result, err := pc.InspectionContract.EvaluateTransaction("QueryAllInspections")
	if err != nil {
		panic(fmt.Errorf("failed to query QueryAllInspections: %w", err))
	}

	fmt.Println("\n*** QueryAllInspections successful")
	fmt.Printf("resultJSON : %s \n", result)

	var resultStruct []Inspection
	if err := json.Unmarshal(result, &resultStruct); err != nil {
		panic(fmt.Errorf("failed to unmarshal result JSON: %s, %w", result, err))
	}

	fmt.Printf("resultStruct %+v", resultStruct)
	return resultStruct
}
