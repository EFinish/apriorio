package main

const (
	templateGenericActive       = `> {{ . | faint | bold }}`
	templateGenericInactive     = `{{ . | faint }}`
	templateSubjectActive       = `> {{ .Body | faint | bold }}`
	templateSubjectInactive     = `{{ .Body | faint }}`
	templatePredicateActive     = `> {{ .Body | faint | bold }}`
	templatePredicateInactive   = `{{ .Body | faint }}`
	templatePremiseActive       = `> {{ . | faint | bold }}`
	templatePremiseInactive     = `{{ . | faint }}`
	templatePropositionActive   = `> {{ . | faint | bold }}`
	templatePropositionInactive = `{{ . | faint }}`
)

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
