package behavioral_patterns

import "fmt"

type Patient struct {
	name string
	registrationDone bool
	doctorVisitDone bool
	medicineDone bool
	paymentDone bool
}

type Department interface {
	execute(*Patient)
	setNext(Department)
}

type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Printf("Registration of patient %v already done\n", p.name)
		r.next.execute(p)
		return
	}

	fmt.Printf("Registering patient %v \n", p.name)
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next;
}

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorVisitDone {
		fmt.Printf("Patient %v already visited the doctor\n", p.name)
		d.next.execute(p)
		return
	}

	fmt.Printf("Doctor is checking patient %v\n", p.name)
	p.doctorVisitDone = true
	d.next.execute(p)
}

func (r *Doctor) setNext(next Department) {
	r.next = next;
}

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Printf("Patient %v already taken medicine", p.name)
		m.next.execute(p)
		return
	}

	fmt.Printf("Patient %v is taking medicine from Medical\n", p.name)
	p.medicineDone = true
	m.next.execute(p)
}

func (r *Medical) setNext(next Department) {
	r.next = next;
}

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Printf("Payment done by patient %v\n", p.name)
		return
	}

	fmt.Printf("Taking payment from patient %v\n", p.name)
	p.paymentDone = true
}

func (r *Cashier) setNext(next Department) {
	r.next = next;
}

func ChainOfResponsibilityStatic() {
	c := &Cashier{}
	
	m := &Medical{}
	m.setNext(c)

	d := &Doctor{}
	d.setNext(m)

	r := &Reception{}
	r.setNext(d)

	p := &Patient{
		name: "John",
	}

	r.execute(p)
}