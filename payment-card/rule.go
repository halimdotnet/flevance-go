package pcard

type Operation int

const (
	Exact Operation = iota
	Range
)

type Rule struct {
	Issuer     IssuerNetwork
	Operation  Operation
	StartValue string
	EndValue   string
	Length     int
}

var mapIssuerRule = []Rule{
	{Issuer: INVisa, Operation: Exact, StartValue: "4", EndValue: ""},
	{Issuer: INAmex, Operation: Exact, StartValue: "34", EndValue: ""},
	{Issuer: INAmex, Operation: Exact, StartValue: "37", EndValue: ""},
	{Issuer: INJCB, Operation: Range, StartValue: "3528", EndValue: "3589"},
	{Issuer: INMastercard, Operation: Range, StartValue: "51", EndValue: "55"},
	{Issuer: INMastercard, Operation: Range, StartValue: "2221", EndValue: "2720"},
	{Issuer: INDiscover, Operation: Exact, StartValue: "65", EndValue: ""},
	{Issuer: INDiscover, Operation: Exact, StartValue: "6011", EndValue: ""},
	{Issuer: INDiscover, Operation: Range, StartValue: "644", EndValue: "649"},
	{Issuer: INDiscover, Operation: Range, StartValue: "622126", EndValue: "622925"},
}
