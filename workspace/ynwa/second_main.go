type Department struct {
    Name         string
    ListEmployees []Employee
}

type Employee struct {
    Name     string
    Position string
    Salary   float64
    Department // анонимное встроенное поле (указатель на Department)
}

func (d *Department) addEmployee(emp Employee) {
    d.ListEmployees = append(d.ListEmployees, emp)
}
