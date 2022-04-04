package ent

import "time"

func (uuo *UserUpdateOne) UpdateUser(input *User, id string) *UserUpdateOne {
	var name string
	if input.Name != name {
		uuo.SetName(input.Name)
	}

	var updated_by string
	if input.UpdatedBy != updated_by {
		uuo.SetUpdatedBy(input.UpdatedBy)
	}

	return uuo
}

func (uc *UserCreate) CreateUser(input *User) *UserCreate {
	var name string
	if input.Name != name {
		uc.SetName(input.Name)
	}
	var update_time time.Time
	if input.UpdateTime != update_time {
		uc.SetUpdateTime(input.UpdateTime)
	}
	var created_time time.Time
	if input.CreatedTime != created_time {
		uc.SetCreatedTime(input.CreatedTime)
	}
	var updated_by string
	if input.UpdatedBy != updated_by {
		uc.SetUpdatedBy(input.UpdatedBy)
	}
	var created_by string
	if input.CreatedBy != created_by {
		uc.SetCreatedBy(input.CreatedBy)
	}
	return uc
}
