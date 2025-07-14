package main

func main() {
	// Create a computer using the builder pattern
	computer := NewComputerBuilder().
		CPU("Intel i7").
		RAM(16).
		MB("ASUS ROG Strix").
		Build()

	println("Computer built with CPU:", computer.CPU, "RAM:", computer.RAM, "MB:", computer.MB)

	// Create an office computer using the office computer builder
	officeComputer := NewOfficeComputerBuilder().Build()
	println("Office Computer built with CPU:", officeComputer.CPU, "RAM:", officeComputer.RAM, "MB:", officeComputer.MB)
}

type Computer struct {
	CPU string
	RAM int
	MB  string
}

type ComputerBuilderI interface {
	CPU(val string) ComputerBuilderI
	RAM(val int) ComputerBuilderI
	MB(val string) ComputerBuilderI

	Build() Computer
}

type computerBuilder struct {
	cpu string
	ram int
	mb  string
}

func NewComputerBuilder() ComputerBuilderI {
	return &computerBuilder{}
}

func (b *computerBuilder) CPU(val string) ComputerBuilderI {
	b.cpu = val
	return b
}

func (b *computerBuilder) RAM(val int) ComputerBuilderI {
	b.ram = val
	return b
}

func (b *computerBuilder) MB(val string) ComputerBuilderI {
	b.mb = val
	return b
}

func (b *computerBuilder) Build() Computer {
	return Computer{
		CPU: b.cpu,
		RAM: b.ram,
		MB:  b.mb,
	}
}

// officeComputerBuilder is just an example how to create a builder with predefined values
type officeComputerBuilder struct {
	computerBuilder
}

func (b officeComputerBuilder) Build() Computer {
	return Computer{
		CPU: b.cpu,
		RAM: b.ram,
		MB:  b.mb,
	}
}

func NewOfficeComputerBuilder() ComputerBuilderI {
	return &officeComputerBuilder{
		computerBuilder: computerBuilder{
			cpu: "Intel i5",
			ram: 8,
			mb:  "ASUS Prime",
		},
	}
}
