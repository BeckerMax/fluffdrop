package helper

type FileHandler interface {
	store(path string, input string)
	retrieve(path string) string
}

func NewlocalFileHandler() *localFileHandler {
	return &localFileHandler{rootpath: "/temp"}
}

type localFileHandler struct {
	rootpath string
}

func (*localFileHandler) store(path string, input string) {
}

func (*localFileHandler) retrieve(path string) string {
	return "nothing"
}
