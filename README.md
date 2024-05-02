# go-invoice

## How to run

```
./goinvoice -customer-name "Google, Inc" ...
```

Below is the full list of arguments

```
-company-cnpj string
  	Your company's CNPJ number (default "12.232.232/0001-22")
-company-name string
  	Your company name (default "Simpsons Software")
-customer-city string
  	The City of the company receiving the invoice (default "Palo Alto")
-customer-name string
  	The name of the company receiving the invoice (default "Google, Inc.")
-customer-state string
  	The State of the company receiving the invoice (default "CA")
-customer-street string
  	The Street Address of the company receiving the invoice (default "Some street in a fancy area, STE 1999")
-customer-zip string
  	The Zip Code of the company receiving the invoice (default "92201")
-hourly-rate float
  	How much are you charging per hour (default 10)
-hours-worked int
  	The amount of hours (qty) to be shown on the invoice line item (default 10)
-rep-cpf string
  	Your company's representative CPF number (default "023.323.323-83")
-rep-name string
  	Your company's representative name (default "Bart Simpson")
-work-desc string
  	Description of the invoice line item: i.e. Software engineering services  (default "Software Engineering Services")
```
