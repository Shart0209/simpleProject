package constants

var Repo struct {
	CreateColums string
	GetColumns   string
}

func init() {
	Repo.CreateColums = "sync_excise_rv"
	Repo.GetColumns = `contract_id,title,contr_number,contr_date,price,start_date,end_date,description,created_at,company_name,company_city,category_name`
}
