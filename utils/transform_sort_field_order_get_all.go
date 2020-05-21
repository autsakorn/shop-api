package utils

import (
	"errors"
	"strings"
)

// TransFormSortFieldOrderGetAll transform sort field string, order string to slice string
// This logic depend on define arguments between controller service storage
// return error if invalid format
func TransFormSortFieldOrderGetAll(
	sortbyStr string, orderStr string,
) (orderby []string, err error) {
	var sortby []string
	if sortbyStr != "" {
		sortby = strings.Split(sortbyStr, ",") // sortbyStr: col1,col2
	}
	var order []string
	if orderStr != "" {
		order = strings.Split(orderStr, ",") // orderStr: desc,asc
	}
	result := []string{}
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				if order[i] == "desc" {
					result = append(result, "-"+v)
				} else if order[i] == "asc" {
					result = append(result, v)
				} else {
					orderby = []string{}
					err = errors.New("Error: Invalid order. Must be either [asc|desc]")
					return
				}
			}
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				if order[0] == "desc" {
					result = append(result, "-"+v)
				} else if order[0] == "asc" {
					result = append(result, v)
				} else {
					orderby = []string{}
					err = errors.New("Error: Invalid order. Must be either [asc|desc]")
					return
				}
			}

		} else if len(sortby) != len(order) && len(order) != 1 {
			err = errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			err = errors.New("Error: unused 'order' fields")
		}
	}
	orderby = result
	return
}
