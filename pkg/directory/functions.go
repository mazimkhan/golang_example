package directory

func AddEmployee(employees map[int]Employee, employee Employee) map[int]Employee {
	result := map[int]Employee{}
	for k, v := range employees {
		result[k] = v
	}
	result[employee.ID] = employee
	return result
}

func RemoveEmployee(employees map[int]Employee, id int) map[int]Employee {
	newMap := map[int]Employee{}
	for k, v := range employees {
		if k != id {
			newMap[k] = v
		}
	}
	return newMap
}

func SearchEmployees(employees map[int]Employee, predicate func(Employee) bool) map[int]Employee {
	var result = map[int]Employee{}
	for _, employee := range employees {
		if predicate(employee) {
			result[employee.ID] = employee
		}
	}
	return result
}
