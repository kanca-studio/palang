package base

type BaseService struct {
	Repository InterfaceBaseRepository
}

func (base *BaseService) Create(params interface{}) interface{} {
	return base.Repository.Create(params)
}

func (base *BaseService) FindById(id uint) (interface{}, error) {
	return base.Repository.FindById(id)
}

func (base *BaseService) Find(filter interface{}) (interface{}, error) {
	return base.Repository.Find(filter)
}

func (base *BaseService) FindAll(filter interface{}) ([]interface{}, error) {
	return base.Repository.FindAll(filter)
}

func (base *BaseService) Update(filter interface{}, param interface{}) error {
	return base.Repository.Update(filter, param)
}

func (base *BaseService) Delete(filter interface{}) error {
	return base.Repository.Delete(filter)

}
// Remove for hard Delete
func (base *BaseService) Remove(filter interface{}) error {
	return base.Repository.Remove(filter)
}