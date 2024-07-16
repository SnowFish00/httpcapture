package packet

import (
	"context"
	"net/http"
	"time"
)

type (
	Event struct {
		Req  *http.Request
		Resp *http.Response
		//test
		Doc map[string]interface{}
	}
	//test
	EventU struct {
		Doc map[string]interface{}
	}
	//test
	EventFunc  func(req *http.Request, resp *http.Response, doc map[string]interface{})
	EventUFunc func(doc map[string]interface{})
	Handle     struct {
		ctx      context.Context // 上下文.
		cardName string          // 网卡名称.
		bpf      string          // 过滤器规则.
		//test
		bpfU    string
		promisc bool             // 是否混杂模式.
		eventCh chan interface{} // 事件通道.
		//test
		eventUCh     chan interface{}
		goroutineNum int       // 协程数量.
		eventHandle  EventFunc // 事件处理.
		//test
		eventHandleU EventUFunc    // 事件处理.
		flushTime    time.Duration // 清理缓存时间.
	}
)

// test
func NewPacketHandle(ctx context.Context, cardName string, eventCh chan interface{}, eventUCh chan interface{}) *Handle {
	return &Handle{
		ctx:       ctx,
		cardName:  cardName,
		bpf:       "tcp",
		promisc:   false,
		eventCh:   eventCh,
		eventUCh:  eventUCh,
		flushTime: time.Minute * -2,
	}
}

// SetBpf 设置过滤器规则.
func (slf *Handle) SetBpf(bpf string) *Handle {
	slf.bpf = bpf
	return slf
}

// test
func (slf *Handle) SetBpfU(bpfU string) *Handle {
	slf.bpfU = bpfU
	return slf
}

// SetPromisc 设置混杂模式.
func (slf *Handle) SetPromisc(promise bool) *Handle {
	slf.promisc = promise
	return slf
}

// SetEventHandle 设置多协程事件处理.
func (slf *Handle) SetEventHandle(goroutineNum int, handle EventFunc) *Handle {
	slf.goroutineNum = goroutineNum
	slf.eventHandle = handle
	return slf
}

// test
func (slf *Handle) SetEventHandleU(goroutineNum int, handleU EventUFunc) *Handle {
	slf.goroutineNum = goroutineNum
	slf.eventHandleU = handleU
	return slf
}

// SetFlushTime 设置清理缓存时间,
// 清除收到的最后一个数据包时间加上此时间之前的所有的数据包.
func (slf *Handle) SetFlushTime(timer time.Duration) *Handle {
	slf.flushTime = timer.Abs() * -1
	return slf
}
