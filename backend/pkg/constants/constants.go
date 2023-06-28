package constants

var Repo struct {
	Columns string
}

func init() {
	Repo.Columns = `title, numb, price, date, start_date, end_date,status,files,created_at,description, 
commons.contract_id, authors.name as author, suppliers.supplier_id as suppliers, categories.category_id as category, c_groups.c_groups_id as сgroups`
}
