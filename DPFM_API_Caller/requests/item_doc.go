package requests

type ItemDoc struct {
	ProductionOrder          int	`json:"ProductionOrder"`
	ProductionOrderItem      int	`json:"ProductionOrderItem"`
	DocType                  string `json:"DocType"`
	FileExtension            string `json:"FileExtension"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
	FilePath                 string `json:"FilePath"`
	FileName                 string `json:"FileName"`
}
