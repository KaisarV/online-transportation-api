package repository

type BaseRepository interface {
	FindByid(id int, model interface{}) (interface{}, error)
	Create(input interface{}) (interface{}, error)
	UpdateByid(id int, input map[string]interface{}, entity interface{}) error
	DeleteByid(id int, entity interface{}) error
}
