package packet

// EventHandle 设置goroutine进行消息处理.
func (slf *Handle) EventHandle() {
	if slf.goroutineNum <= 0 {
		return
	}

	for i := 0; i < slf.goroutineNum; i++ {
		go slf.handle()
	}

	select {
	case <-slf.ctx.Done():
		close(slf.eventCh)
		return
	}

}

// test
func (slf *Handle) EventHandleU() {
	if slf.goroutineNum <= 0 {
		return
	}

	for i := 0; i < slf.goroutineNum; i++ {
		go slf.handleU()
	}

	select {
	case <-slf.ctx.Done():
		close(slf.eventUCh)
		return
	}

}

// handle 消息处理.
func (slf *Handle) handle() {
	for val := range slf.eventCh {
		data := val.(Event)
		slf.eventHandle(data.Req, data.Resp, data.Doc)
	}
}

// test
func (slf *Handle) handleU() {
	for val := range slf.eventUCh {
		data := val.(EventU)
		slf.eventHandleU(data.Doc)
	}
}
