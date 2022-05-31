package config

type Config struct {
	PathLength                  int      `json:"pathLength"`                  // 数据集数目
	DataPaths                   []string `json:"dataFilePaths"`               // 数据文件集
	DataIndexPaths              []string `json:"dafaIndexPaths"`              // 数据索引文件集
	CsvPaths                    []string `json:"csvPaths"`                    // csv文件集
	InvertedIndexFilePath       string   `json:"invertedIndexFilePath"`       // 倒排索引文件
	DictPath                    string   `json:"dictPath"`                    // 词典文件
	StopWordPath                string   `json:"stopWordPath"`                // 停用词表
	ImportCsvCoroutines         int      `json:"importCsvCoroutines"`         // 导入csv文件时的并发数
	MakeInvertedIndexCoroutines int      `json:"makeInvertedIndexCoroutines"` // 构造倒排索引的并发数
	SearchLruMaxCapacity        int      `json:"searchLruMaxCapacity"`        // 最外层搜索的lru缓存大小
}
