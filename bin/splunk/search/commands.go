package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

const DATA_DICT = 0
const DATA_CSVROW = 1
const DATA_RAW = 2

var headerRe = regexp.MustCompile(`chunked\s+1.0,(?P<metadata_length>\d+),(?P<body_length>\d+)`)

type Metadata map[string]string
type Data []map[string]string

type SearchCommand interface {
	HandleGetInfo(metadata Metadata) Metadata
	HandleExecute(metadata Metadata, data Data) Data
	HandleResults(metadata Metadata, data Data) Data
}


type Executor struct {
	Reader  *bufio.Reader
	Writer  io.Writer
	Command SearchCommand
	HandlerData int
}

func NewExecutor(reader io.Reader, writer io.Writer, command SearchCommand) *Executor {
	executor := new(Executor)
	executor.Reader = bufio.NewReader(reader)
	executor.Writer = writer
	executor.Command = command
	return executor
}

func (self *Executor) Run() (err error) {
	for err == nil {
		err := self._handleChunk()
	}
	if self.HandlerData = DATA_DICT {

	}
	elseif self.HandlerData = DATA_CSVROW {
		
	}

	return
}

func (self *Executor) _handleChunk() (err error) {
	metadata, body, err := self._readChunk()
	if err != nil {
		return
	}

}

func (self *Executor) _readChunk() (metadata map[string]interface{}, body string, err error) {
	header, err := self.Reader.ReadString('\n')
	if len(header) == 0 {
		metadata = nil
		return
	}

	matches := headerRe.FindAllStringSubmatch(header, -1)
	if len(matches) == 0 {
		err = fmt.Errorf("Failed to parse transport header: %s", header)
		return
	}

	metadataLength, err := strconv.ParseInt(matches[0][1], 10, 32)
	bodyLength, err := strconv.ParseInt(matches[0][2], 10, 32)

	metadataBuffer := make([]byte, metadataLength)
	self.Reader.Read(metadataBuffer)

	bodyBuffer := make([]byte, bodyLength)
	self.Reader.Read(bodyBuffer)

	body = string(bodyBuffer)
	err = json.Unmarshal(metadataBuffer, &metadata)
	return
}
