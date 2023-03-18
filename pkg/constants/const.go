package constants

var Repo struct {
	CreateColums string
	GetColumns   string
}

func init() {
	Repo.CreateColums = `title, contr_number, contr_date, category_id, price, start_date, end_date,description, distributor_id, files`
	Repo.GetColumns = `contract_id,title,contr_number,contr_date,price,start_date,end_date,description,created_at,company_name,category_name,files`
}
