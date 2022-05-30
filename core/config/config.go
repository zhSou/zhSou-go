package config

type Config struct {
	DataPaths             []string `json:"dataFilePaths"`         // 数据文件集
	DataIndexPaths        []string `json:"dafaIndexPaths"`        // 数据索引文件集
	CsvPaths              []string `json:"csvPaths"`              // csv文件集
	InvertedIndexFilePath string   `json:"invertedIndexFilePath"` // 倒排索引文件
	StopWordPath          string   `json:"stopWordPath"`          // 停用词表
}
