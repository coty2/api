package domain

import(
	v "suffgo/internal/option/domain/valueObjects"
)

type(
	Option struct {
		id *v.ID
		value *v.Value
	}

	OptionDTO struct {
		ID uint `json:"id"`
		Value string `json:"value"`
	}

	OptionCreateRequest struct {
		Value string `json:"value`
	}
)

func NewOption(id *v.ID, value v.Value)
	*Option{
		return &Option{
			id:	id,
			value: value,
		}
	}

func(o *Option) ID() v.ID{
	return *o.id
}

func(o *Option) Value() v.Value{
	return o.value
}