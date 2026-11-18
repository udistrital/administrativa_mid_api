package models

type RpDocente struct {
	CdpRpDocente CdpRpDocente `json:"cdp_rp_docente"`
}

type CdpRpDocente struct {
	CdpRp []CdpRp `json:"cdp_rp"`
}

type CdpRp struct {
	Cdp      string `json:"cdp"`
	Vigencia string `json:"vigencia"`
	Rp       string `json:"rp"`
}
