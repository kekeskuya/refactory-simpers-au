package constants

const (
	MessageSuccess = "Success"

	DataNPWP   = "npwp"
	DataAsabri = "asabri"
	DataPaspor = "paspor"
)

type Category string

const (
	CategoryNPWP   Category = DataNPWP
	CategoryAsabri Category = DataAsabri
	CategoryPaspor Category = DataPaspor
)
