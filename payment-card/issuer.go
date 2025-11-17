package pcard

type IssuerNetwork string

const (
	INMastercard IssuerNetwork = "Mastercard"
	INVisa                     = "Visa"
	INDiscover                 = "Discover"
	INAmex                     = "American Express"
	INJCB                      = "JCB"
)

func getIssuer(input string) IssuerNetwork {
	for _, rule := range mapIssuerRule {

		l := len(rule.StartValue)
		p := input[0:l]

		switch rule.Operation {
		case Exact:
			if p == rule.StartValue {
				return rule.Issuer
			}
		case Range:
			if p >= rule.StartValue && p <= rule.EndValue {
				return rule.Issuer
			}
		}
	}

	return ""
}
