package factor3

var log Logger

func init() {
	log = NewLogger()

	log.SetLevel(InfoLevel)
	log.Info("initializing logger", nil)
}
