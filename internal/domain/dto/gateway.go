package dto

type LoginRequest struct {
	UserType                   string `json:"userType"`
	Email                      string `json:"email"`
	Name                       string `json:"name"`
	SnsProvider                string `json:"snsProvider"`
	Image                      string `json:"img"`
	ExternalId                 string `json:"externalId"`
	SocialProvidedAccessToken  string `json:"socialProvidedAccessToken"`
	SocialProvidedRefreshToken string `json:"socialProvidedRefreshToken"`
}

type LoginResponse struct {
	UserPayload   *UserPayload   `json:"userPayload"`
	AccessToken   string         `json:"accessToken"`
	RefreshToken  string         `json:"refreshToken"`
	RefreshExpire uint64         `json:"refreshExpire"`
	HederaAccount *HederaAccount `json:"hederaAccount"`
}

type HederaAccount struct {
	EvmAddress string `json:"evmAddress"`
}

type UserPayload struct {
	Id               string `json:"id"`
	Uuid             uint64 `json:"uuid"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	SnsProvider      string `json:"snsProvider"`
	HederaEvmAddress string `json:"hederaEvmAddress"`
}

type MintDHNRequest struct {
	Metadata    *Metadata `json:"metadata"`
	MetadataUri string    `json:"metadataUri"`
	ImageUri    string    `json:"imageUri"`
	SegmentId   string    `json:"segmentId"`
}

type MintDHNResponse struct {
	Result Result `json:"receipt"`
}

type Metadata struct {
	Image      string   `json:"image"`
	Attributes []string `json:"attributes"`
}

type Result struct {
	Message string  `json:"message"`
	Data    Data    `json:"data"`
	Receipt Receipt `json:"receipt"`
}

type Receipt struct {
	BlockHash         string `json:"blockHash"`
	BlockNumber       string `json:"blockNumber"`
	From              string `json:"from"`
	To                string `json:"to"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	ContractAddress   string `json:"contractAddress"`
	Logs              []Log  `json:"logs"`
	LogsBloom         string `json:"logsBloom"`
	TransactionHash   string `json:"transactionHash"`
	TransactionIndex  string `json:"transactionIndex"`
	EffectiveGasPrice string `json:"effectiveGasPrice"`
	Status            string `json:"status"`
	Type              string `json:"type"`
}

type Log struct {
	Address          string   `json:"address"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Data             string   `json:"data"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

type Data struct {
	TokenId uint64 `json:"tokenId"`
}
