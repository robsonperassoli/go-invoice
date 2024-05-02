package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type Invoice struct {
	Customer Customer
	Company  CompanyDetails
	LineItem LineItem
}

type Customer struct {
	Name          string
	StreetAddress string
	City          string
	State         string
	ZipCode       string
}

type CompanyDetails struct {
	CompanyName        string
	CompanyCNPJ        string
	RepresentativeName string
	RepresentativeCPF  string
}

type LineItem struct {
	Description string
	UnitPrice   float64
	Quantity    int64
	Total       float64
}

func parseInput() Invoice {
	custName := flag.String("customer-name", "Google, Inc.", "The name of the company receiving the invoice")
	custStreet := flag.String("customer-street", "Some street in a fancy area, STE 1999", "The Street Address of the company receiving the invoice")
	custCity := flag.String("customer-city", "Palo Alto", "The City of the company receiving the invoice")
	custState := flag.String("customer-state", "CA", "The State of the company receiving the invoice")
	custZip := flag.String("customer-zip", "92201", "The Zip Code of the company receiving the invoice")

	workDesc := flag.String("work-desc", "Software Engineering Services", "Description of the invoice line item: i.e. Software engineering services ")
	hoursWorked := flag.Int64("hours-worked", 10, "The amount of hours (qty) to be shown on the invoice line item")
	hourlyRate := flag.Float64("hourly-rate", 10.0, "How much are you charging per hour")

	compName := flag.String("company-name", "Simpsons Software", "Your company name")
	compCnpj := flag.String("company-cnpj", "12.232.232/0001-22", "Your company's CNPJ number")
	repName := flag.String("rep-name", "Bart Simpson", "Your company's representative name")
	repCpf := flag.String("rep-cpf", "023.323.323-83", "Your company's representative CPF number")

	flag.Parse()

	customer := Customer{
		Name:          *custName,
		StreetAddress: *custStreet,
		City:          *custCity,
		State:         *custState,
		ZipCode:       *custZip,
	}

	total := *hourlyRate * float64(*hoursWorked)
	lineItem := LineItem{
		Description: *workDesc,
		Quantity:    *hoursWorked,
		UnitPrice:   *hourlyRate,
		Total:       total,
	}

	company := CompanyDetails{
		CompanyName:        *compName,
		CompanyCNPJ:        *compCnpj,
		RepresentativeName: *repName,
		RepresentativeCPF:  *repCpf,
	}

	return Invoice{
		Company:  company,
		Customer: customer,
		LineItem: lineItem,
	}
}

func renderPdf(invoice Invoice, date string, filePath string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(20, 30, 20)
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// DOC TITLE
	pdf.Write(10, "INVOICE")
	pdf.Ln(13)

	pdf.SetFontStyle("U")
	pdf.SetFontSize(7)
	pdf.Cell(0, 7, "Prepared for")
	pdf.Ln(4)

	pdf.SetFontStyle("")
	pdf.SetFontSize(8)

	pdf.Write(10, "Company Name: ")
	pdf.SetFontStyle("B")
	pdf.Write(10, invoice.Customer.Name)
	pdf.SetFontStyle("")
	pdf.Ln(4)

	pdf.Write(10, "Company Address: ")
	pdf.SetFontStyle("B")
	pdf.Write(10, invoice.Customer.StreetAddress)
	pdf.SetFontStyle("")
	pdf.Ln(4)

	pdf.Write(10, "City, State, Zip: ")
	pdf.SetFontStyle("B")
	pdf.Write(10, fmt.Sprintf("%s, %s %s", invoice.Customer.City, invoice.Customer.State, invoice.Customer.ZipCode))
	pdf.SetFontStyle("")
	pdf.Ln(20)

	// ITEMS TABLE HEADER
	pdf.SetFont("Arial", "", 9)
	pdf.SetFillColor(240, 240, 240)
	pdf.CellFormat(90, 8, "DESCRIPTION OF WORK", "1", 0, "L", true, 0, "")
	pdf.CellFormat(25, 8, "QTY/HRS", "1", 0, "L", true, 0, "")
	pdf.CellFormat(25, 8, "UNIT PRICE", "1", 0, "L", true, 0, "")
	pdf.CellFormat(30, 8, "SUBTOTAL", "1", 1, "L", true, 0, "")

	// ITEMS TABLE ROWS
	pdf.SetFillColor(255, 255, 255)
	pdf.CellFormat(90, 8, invoice.LineItem.Description, "1", 0, "L", false, 0, "")
	pdf.CellFormat(25, 8, fmt.Sprintf("%d hrs", invoice.LineItem.Quantity), "1", 0, "R", false, 0, "")
	pdf.CellFormat(25, 8, fmt.Sprintf("$%f/hr", invoice.LineItem.UnitPrice), "1", 0, "R", false, 0, "")
	pdf.CellFormat(30, 8, fmt.Sprintf("$%f", invoice.LineItem.Total), "1", 1, "R", false, 0, "")
	pdf.Ln(20)

	// TOTALS
	pdf.SetX(130)
	pdf.SetFillColor(240, 240, 240)
	pdf.SetFontStyle("B")
	pdf.CellFormat(30, 8, "GRAND TOTAL", "1", 0, "C", true, 0, "")
	pdf.SetFontStyle("")
	pdf.CellFormat(30, 8, fmt.Sprintf("$%f", invoice.LineItem.Total), "1", 0, "R", true, 0, "")

	// PAYMENT TERMS
	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(20, 135, 170, 45, "F")

	pdf.SetXY(25, 138)
	pdf.SetFontStyle("B")
	pdf.Cell(40, 10, "PAYMENT TERMS")
	pdf.Ln(7)

	pdf.SetFontStyle("")
	pdf.SetFontSize(8)
	pdf.SetX(25)
	pdf.Cell(40, 10, "To be made payable to")
	pdf.Ln(6)

	pdf.SetFontStyle("B")
	pdf.SetX(25)
	pdf.Cell(40, 10, invoice.Company.RepresentativeName)
	pdf.Ln(4)
	pdf.SetX(25)
	pdf.Cell(40, 10, fmt.Sprintf("CPF: %s", invoice.Company.RepresentativeCPF))
	pdf.Ln(7)

	pdf.SetX(25)
	pdf.Cell(0, 10, invoice.Company.CompanyName)
	pdf.Ln(4)
	pdf.SetX(25)
	pdf.Cell(0, 10, fmt.Sprintf("CNPJ: %s", invoice.Company.CompanyCNPJ))
	pdf.Ln(4)

	pdf.SetXY(140, 138)
	pdf.SetFontSize(10)
	pdf.Cell(0, 10, fmt.Sprintf("DATE: %s", date))

	return pdf.OutputFileAndClose(filePath)
}

func main() {
	var invoice Invoice = parseInput()
	date := time.Now().Format("Mon, Jan _2, 2006")

	err := renderPdf(invoice, date, "invoice.pdf")
	if err != nil {
		fmt.Errorf("error generating pdf", err)
	}
}
