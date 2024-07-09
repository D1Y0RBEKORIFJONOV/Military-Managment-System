package storehouse

func (repo *StorehouseRepository) Returning() string {
	return "RETURNING RETURNING id, name, price amount, type_artillery, created_at, updated_at, deleted_at"
}

func (reop *StorehouseRepository) SelectQuery() string {
	return `
	id,
	name,
	price
	amount,
	type_artillery,
	created_at,
	updated_at,
	deleted_at`
}
