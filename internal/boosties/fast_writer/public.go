package fast_writer

func (writer *Writer) Write(contentType string, status int, body []byte) {
	writer.mutex.Lock()
	defer writer.mutex.Unlock()

	if writer.wasWritten {
		return
	}

	writer.request.SetStatusCode(status)
	writer.request.Response.Header.Set("Content-Type", contentType)
	writer.request.SetBody(body)

	writer.wasWritten = true
}
