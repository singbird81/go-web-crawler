package downloader

import "errs"

// genError 用于生成爬虫错误值。
func genError(errMsg string) error {
	return errs.NewCrawlerError(errs.ERROR_TYPE_DOWNLOADER,
		errMsg)
}

// genParameterError 用于生成爬虫参数错误值。
func genParameterError(errMsg string) error {
	return errs.NewCrawlerErrorBy(errs.ERROR_TYPE_DOWNLOADER,
		errs.NewIllegalParameterError(errMsg))
}
