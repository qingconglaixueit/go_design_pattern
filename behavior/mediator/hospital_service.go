// @Author Bing 
// @Desc
package main

type HosService struct {
	IsFree bool
	Queue  []IPeople
}

func (h *HosService) NoticePeople() {
	if !h.IsFree {
		h.IsFree = true
		// 如果队列有人等待，取出第一个准入
		if len(h.Queue) > 0 {
			q := h.Queue[0]
			h.Queue = h.Queue[1:]
			q.PermitEnter()
		}
	}
}

func (h *HosService) CanEnter(people IPeople) bool {
	if h.IsFree {
		h.IsFree = false
		return true
	}

	h.Queue = append(h.Queue, people)
	return false
}
