package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1000", "Rangga", "Senayan City", "00987", "1999-08-25", "1"},
		{"1001", "Gilang", "Bandung", "04647", "1999-08-05", "1"},
		{"1002", "Samid", "Tegal", "882625", "1999-08-20", "1"},
		{"1003", "Lutung", "Brebes", "120987", "1998-08-21", "1"},
	}
	return CustomerRepositoryStub{customers}
}
