package model

type CoverageModel struct {
	Id    int64  `json:"id"`
	Kdcab string `json:"kdcab"`
	Nik   string `json:"nik"`
}

type CoverageCreatUpdateModel struct {
	Id    int64  `json:"id" validate:"required"`
	Kdcab string `json:"kdcab" validate:"required"`
	Nik   string `json:"nik" validate:"required"`
}
