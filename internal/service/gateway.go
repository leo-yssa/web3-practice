package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"web3-practice/internal/domain/dto"
)

type GatewayService interface {
	Login(gatewayLoginRequest *dto.LoginRequest) (*dto.LoginResponse, error)
	MintDHN(mintDHNRequest *dto.MintDHNRequest, accessToken string) (*dto.MintDHNResponse, error)
}

func NewGatewayService(host string) GatewayService {
	return &gatewayService{
		Client: &http.Client{},
		host:   host,
	}
}

type gatewayService struct {
	*http.Client
	host string
}

func (gs *gatewayService) post(path string, request interface{}, accessToken string) ([]byte, error) {
	rb, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest(http.MethodPost, gs.host+path, bytes.NewBuffer(rb))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")
	if accessToken != "" {
		r.Header.Add("Authorization", "Bearer "+accessToken)
	}
	response, err := gs.Do(r)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

func (gs *gatewayService) Login(loginRequest *dto.LoginRequest) (*dto.LoginResponse, error) {
	body, err := gs.post("/login", loginRequest, "")
	if err != nil {
		return nil, err
	}
	response := &dto.LoginResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (gs *gatewayService) MintDHN(mintDHNRequest *dto.MintDHNRequest, accessToken string) (*dto.MintDHNResponse, error) {
	body, err := gs.post("/digitalhumannft/mint", mintDHNRequest, accessToken)
	if err != nil {
		return nil, err
	}
	mintDHNResponse := &dto.MintDHNResponse{}
	err = json.Unmarshal(body, mintDHNResponse)
	if err != nil {
		return nil, err
	}
	return mintDHNResponse, nil
}

// func (gs *gatewayService) MintRootInventory(gatewayRootInventory *dto.GatewayRootInventoryRequest) (*dto.GatewayRootInventoryResponse, error) {
// 	b, err := json.Marshal(gatewayRootInventory)
// 	if err != nil {
// 		return nil, err
// 	}
// 	buff := bytes.NewBuffer(b)
// 	r, err := http.NewRequest(http.MethodPost, gs.host+"/inventory/root-inventory", buff)
// 	if err != nil {
// 		return nil, err
// 	}
// 	r.Header.Add("Authorization", "Bearer "+gatewayRootInventory.AccessToken)
// 	c := &http.Client{}
// 	resp, err := c.Do(r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	gatewayRootInventoryResponse := &dto.GatewayRootInventoryResponse{}
// 	err = json.Unmarshal(body, gatewayRootInventoryResponse)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return gatewayRootInventoryResponse, nil
// }

// func (gs *gatewayService) MintInventory(gatewayInventory *dto.GatewayInventoryRequest) (*dto.GatewayInventoryResponse, error) {
// 	b, err := json.Marshal(gatewayInventory)
// 	if err != nil {
// 		return nil, err
// 	}
// 	buff := bytes.NewBuffer(b)
// 	r, err := http.NewRequest(http.MethodPost, gs.host+"/inventory/inventory", buff)
// 	if err != nil {
// 		return nil, err
// 	}
// 	r.Header.Add("Authorization", "Bearer "+gatewayInventory.AccessToken)
// 	c := &http.Client{}
// 	resp, err := c.Do(r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	gatewayInventoryResponse := &dto.GatewayInventoryResponse{}
// 	err = json.Unmarshal(body, gatewayInventoryResponse)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return gatewayInventoryResponse, nil
// }

// func (gs *gatewayService) CreateCampaign(gatewayCampaign *dto.GatewayCampaignRequest) (*dto.GatewayCampaignResponse, error) {
// 	b, err := json.Marshal(gatewayCampaign)
// 	if err != nil {
// 		return nil, err
// 	}
// 	buff := bytes.NewBuffer(b)
// 	r, err := http.NewRequest(http.MethodPost, gs.host+"/factory/campaign", buff)
// 	if err != nil {
// 		return nil, err
// 	}
// 	r.Header.Add("Authorization", "Bearer "+gatewayCampaign.AccessToken)
// 	c := &http.Client{}
// 	resp, err := c.Do(r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	gatewayCampaignResponse := &dto.GatewayCampaignResponse{}
// 	err = json.Unmarshal(body, gatewayCampaignResponse)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return gatewayCampaignResponse, nil
// }
