package command

type BufferWriter struct {
	history []string
}

func NewBufferWriter() *BufferWriter {
	return &BufferWriter{
		history: []string{},
	}
}

func (bufferWriter *BufferWriter) Write(message string) error {
	bufferWriter.history = append(bufferWriter.history, message)
	return nil
}

func (bufferWriter *BufferWriter) GetHistory() []string {
	return bufferWriter.history
}
