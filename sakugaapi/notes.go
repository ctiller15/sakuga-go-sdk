package sakugaapi

type NotesAPI struct {
	URL string
}

func newNotesAPI(baseURL string) *NotesAPI {
	newAPI := NotesAPI{
		URL: baseURL + "/note",
	}
	return &newAPI
}
