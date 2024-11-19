package usecases

var FacebookEntry FacebookEntryInterface

type FacebookEntryInterface interface {
	Authentication()
	GetCallFb()
}

type facebookEntry struct {
}

func (f *facebookEntry) Authentication() {
	// Implement the method
}

func (f *facebookEntry) GetCallFb() {
	// Implement the method
}

func NewFacebookEntry() FacebookEntryInterface {
	FacebookEntry = &facebookEntry{}
	return FacebookEntry
}
