package constants

var Repo struct {
	Colum string
}

func init() {
	Repo.Colum = `title, contr_number, contr_date, price, start_date, end_date, description`
}
