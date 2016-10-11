package plugin

import (
	"github.com/bakhi/test"
	"github.com/bakhi/udsfs"
	"gopkg.in/sensorbee/sensorbee.v0/bql"
)

func init() {
	bql.MustRegisterGlobalSourceCreator("lorem", bql.SourceCreatorFunc(udsfs.CreateLoremSource))
	bql.MustRegisterGlobalSourceCreator("src", bql.SourceCreatorFunc(test.CreateMySource))
	//	bql.MustRegisterGlobalSourceCreator("lorem", bql.SourceCreatorFunc(udsfs.CreateLoremSource))
	//	udf.MustRegisterGlobalUDSFCreator("word_splitter", udf.MustConvertToUDSFCreator(udsfs.CreateWordSplitter))
	//	udf.MustRegisterGlobalUDSFCreator("ticker", udf.MustConvertToUDSFCreator(udsfs.CreateTicker))
}
