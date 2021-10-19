package staff

import "fmt"

type Contract struct {
	empId    int
	basicpay int
}

func (c *Contract) Print() {
	fmt.Printf(" Mã NV: %v   Lương: %10v\n", c.empId, c.basicpay)
}

func (c *Contract) CalculateSalary() int {
	return c.basicpay
}

func (c *Contract) SetEmpID(empId int) {
	c.empId = empId
}

func (c *Contract) SetBasicPay(basicpay int) {
	c.basicpay = basicpay
}

func (c *Contract) SetPF(pf int) {
}

func (c *Contract) Clone() iStaff {
	return &Contract{
		empId:    c.empId,
		basicpay: c.basicpay,
	}
}
