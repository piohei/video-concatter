package downloader

type Downloader interface {
	Download(url, target string) error

	setRetries(n int)
}

func Retries(n int) func(Downloader) {
	return func(downloader Downloader) {
		downloader.setRetries(n)
	}
}

func NewDownloader(options ...func(Downloader)) Downloader {
	d := &defaultDownloader{}
	for _, o := range options {
		o(d)
	}
	return d
}
