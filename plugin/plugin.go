package plugin

import (
	"github.com/bakhi/test"
	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
)

func init() {
	bql.MustRegisterGlobalSourceCreator("my_src", bql.SourceCreatorFunc(test.CreateMySource))
	//	bql.MustRegisterGlobalSourceCreator("lorem", bql.SourceCreatorFunc(udsfs.CreateLoremSource))
	//	udf.MustRegisterGlobalUDSFCreator("word_splitter", udf.MustConvertToUDSFCreator(udsfs.CreateWordSplitter))
	udf.MustRegisterGlobalUDSFCreator("ticker", udf.MustConvertToUDSFCreator(test.CreateTicker))
}
