package dataset

import (
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhSou/zhSou-go/util/algorithm/binary"
	"io"
	"log"
	"os"
)

// SeekInfo 存放某条记录的寻址信息
type SeekInfo struct {
	Offset     uint32
	UrlLength  uint16
	TextLength uint16
}

// IndexFile 存放索引文件信息
type IndexFile struct {
	ItemLength uint32
	SeekInfo   []SeekInfo
}

func (f *IndexFile) SaveToFile(path string) error {
	outputFile, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
	defer outputFile.Close()
	enc := gob.NewEncoder(outputFile)
	if err := enc.Encode(*f); err != nil {
		log.Fatalf("Gob文件写入失败  path %s  err %v", path, err)
		return err
	}
	log.Println("Gob文件成功输出：" + path)
	return nil
}

func ConvCsvMakeIndexFile(inputCsvPath string, outputDataPath string, outputIndexPath string) {
	file, err := os.Open(inputCsvPath)
	if err != nil {
		log.Printf("csv文件打开失败 path %s err %v", inputCsvPath, err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	log.Printf("打开成功 path %s ", inputCsvPath)

	outputFile, err := os.OpenFile(outputDataPath, os.O_RDWR|os.O_CREATE, 0777)
	defer outputFile.Close()

	contentInfo := IndexFile{
		ItemLength: 0,
		SeekInfo:   []SeekInfo{},
	}

	// 字节偏移量
	var offset uint32 = 0
	var outputBytes []byte
	for i := 0; ; i++ {
		cols, err := reader.Read()
		if err == io.EOF {
			fmt.Println(i)
			break
		}
		if err != nil {
			log.Fatalln("csv文件读取失败：", err)
		}
		// 跳过csv表头
		if i == 0 {
			continue
		}

		bs1 := []byte(cols[0])
		bs2 := []byte(cols[1])

		urlLen := uint16(len(bs1))
		textLen := uint16(len(bs2))
		contentInfo.SeekInfo = append(
			contentInfo.SeekInfo,
			SeekInfo{
				offset,
				urlLen,
				textLen,
			})
		outputBytes = append(outputBytes, bs1...)
		outputBytes = append(outputBytes, bs2...)

		offset += uint32(urlLen + textLen)
		contentInfo.ItemLength++
	}
	outputFile.Write(outputBytes)
	outputFile.Sync()
	contentInfo.SaveToFile(outputIndexPath)
}

// IndexFileSet 数据集索引
type IndexFileSet struct {
	indexFiles []IndexFile

	// 这个字段为indexFile中itemLength的前缀和
	idArray []uint32
}

// NewIndexFileSet  从文件内容来构造数据集
func NewIndexFileSet(rs []io.Reader) *IndexFileSet {
	var indexFileSet IndexFileSet
	for i, r := range rs {
		var indexFile IndexFile
		gob.NewDecoder(r).Decode(&indexFile)
		indexFileSet.indexFiles = append(indexFileSet.indexFiles, indexFile)
		indexFileSet.idArray = append(indexFileSet.idArray, indexFile.ItemLength)
		if i > 0 {
			indexFileSet.idArray[i] += indexFileSet.idArray[i-1]
		}
	}
	return &indexFileSet
}

func (i *IndexFileSet) Get(id uint32) (fileId int, seekInfo *SeekInfo) {
	nsa := binary.NewSliceAccessible[uint32](i.idArray)
	fileId = binary.FindFirstBigger[uint32](nsa, id)

	var innerSeekInfoId = id
	if fileId != 0 {
		innerSeekInfoId = id - i.idArray[fileId-1]
	}

	seekInfo = &i.indexFiles[fileId].SeekInfo[innerSeekInfoId]
	return
}

type DataReader struct {
	indexFilePaths []string
	dataFilePaths  []string
	indexFileSet   *IndexFileSet
	fileCache      []io.ReaderAt
}

func NewDataReader(indexFilePaths []string, dataFilePaths []string) (*DataReader, error) {
	if len(indexFilePaths) != len(dataFilePaths) {
		return nil, errors.New("索引文件集与数据文件集数量不一致")
	}
	var indexFileSetReaders []io.Reader
	var fileCache []io.ReaderAt

	for _, path := range indexFilePaths {
		indexFile, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		indexFileSetReaders = append(indexFileSetReaders, indexFile)
		fileCache = append(fileCache, nil)
	}
	return &DataReader{
		indexFilePaths: indexFilePaths,
		dataFilePaths:  dataFilePaths,
		indexFileSet:   NewIndexFileSet(indexFileSetReaders),
		fileCache:      fileCache,
	}, nil
}

func (r *DataReader) getReaderAt(fileId int) (io.ReaderAt, error) {
	if r.fileCache[fileId] == nil {
		filePath := r.dataFilePaths[fileId]
		log.Println("尝试加载数据文件: ", filePath)
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		r.fileCache[fileId] = file
	}
	return r.fileCache[fileId], nil
}

type DataRecord struct {
	Url  string
	Text string
}

func (r *DataReader) Read(id uint32) (*DataRecord, error) {
	fileId, seekInfo := r.indexFileSet.Get(id)
	fileReader, err := r.getReaderAt(fileId)
	if err != nil {
		return nil, err
	}
	urlBs := make([]byte, seekInfo.UrlLength)
	textBs := make([]byte, seekInfo.TextLength)
	_, _ = fileReader.ReadAt(urlBs, int64(seekInfo.Offset))
	_, _ = fileReader.ReadAt(textBs, int64(seekInfo.Offset)+int64(seekInfo.UrlLength))
	return &DataRecord{
		string(urlBs),
		string(textBs),
	}, nil
}

func (r *DataReader) Len() uint32 {
	idA := r.indexFileSet.idArray
	return idA[len(idA)-1]
}
