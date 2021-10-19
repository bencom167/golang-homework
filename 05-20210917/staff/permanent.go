package staff

import "fmt"

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

func (p *Permanent) Print() {
	fmt.Printf(" Mã NV: %v   Lương: %10v   Thưởng: %10v\n", p.empId, p.basicpay, p.pf)
}

func (p *Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (p *Permanent) SetEmpID(empId int) {
	p.empId = empId
}

func (p *Permanent) SetBasicPay(basicpay int) {
	p.basicpay = basicpay
}

func (p *Permanent) SetPF(pf int) {
	p.pf = pf
}

func (p *Permanent) Clone() iStaff {
	return &Permanent{
		empId:    p.empId,
		basicpay: p.basicpay,
		pf:       p.pf,
	}
}
