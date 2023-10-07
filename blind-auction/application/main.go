package main

import (
    "fmt"
    "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
    "github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
    "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
    "github.com/hyperledger/fabric-sdk-go/pkg/fab"
    "github.com/hyperledger/fabric-sdk-go/pkg/fab/resource"
    "github.com/hyperledger/fabric-sdk-go/pkg/fab/mocks"
    "github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
    "github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
)


		const (
			channelID      = "mychannel"    
			configFile     = "config.yaml"
			orgName        = "Org1"        
			userName       = "Admin"       
			chaincodeID    = "blind_auction"
	)
	
	func main() {
			
			sdk, err := fabsdk.New(config.FromFile(configFile))
			if err != nil {
					fmt.Printf("Error while creating the SDK: %s\n", err)
					return
			}
			defer sdk.Close()
	
			
			org1ChannelClientContext := sdk.ChannelContext(channelID, fabsdk.WithUser(userName), fabsdk.WithOrg(orgName))
	
			
			client, err := channel.New(org1ChannelClientContext)
			if err != nil {
					fmt.Printf("Erro ao criar o cliente de canal: %s\n", err)
					return
			}
	
			
			chaincode := channel.Request{ChaincodeID: chaincodeID}
	
			
			response, err := client.Execute(chaincode, "CreateAuction", "auction1", "Descrição do Leilão 1")
			if err != nil {
					fmt.Printf("Erro ao criar o leilão: %s\n", err)
					return
			}
			fmt.Printf("Resposta da criação do leilão: %s\n", response.TransactionID)
	
			
			response, err = client.Execute(chaincode, "PlaceBid", "auction1", "Bidder1", "100")
			if err != nil {
					fmt.Printf("Erro ao fazer um lance no leilão: %s\n", err)
					return
			}
			fmt.Printf("Resposta do lance: %s\n", response.TransactionID)
	
			
			response, err = client.Execute(chaincode, "RevealBid", "auction1", "Bidder1", "100")
			if err != nil {
					fmt.Printf("Erro ao revelar o lance: %s\n", err)
					return
			}
			fmt.Printf("Resposta da revelação do lance: %s\n", response.TransactionID)
	
			
			response, err = client.Execute(chaincode, "CloseAuction", "auction1")
			if err != nil {
					fmt.Printf("Erro ao fechar o leilão: %s\n", err)
					return
			}
			fmt.Printf("Resposta do fechamento do leilão: %s\n", response.TransactionID)
	}
