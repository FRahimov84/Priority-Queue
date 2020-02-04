package main

type pQueueMethods interface {
	len() int
	first() *pElement
	last() *pElement
	equeue(value interface{}, priority int)
	dequeue() interface{}
	makeQueueOfPriotiry()
}

type pElement struct {
	previous *pElement
	next     *pElement
	value    interface{}
	priority int
}

type pQueue struct {
	front *pElement
	back  *pElement
	size  int
}

func (receiver *pQueue) makeQueueOfPriotiry() {
	i := receiver.first()
	for i != nil {
		j := receiver.first()
		for j != nil {
			if i.priority < j.priority {
				i.value, j.value = j.value, i.value
				i.priority, j.priority = j.priority, i.priority
			}
			j = j.next
		}
		i = i.next
	}
}

func (receiver *pQueue) len() int {
	return receiver.size
}

func (receiver *pQueue) first() *pElement {
	return receiver.front
}

func (receiver pQueue) last() *pElement {
	return receiver.back
}

func (receiver *pQueue) equeue(value interface{}, priority int) {
	if receiver.size == 0 {
		receiver.size++
		receiver.front = &pElement{
			previous: nil,
			next:     nil,
			value:    value,
			priority: priority,
		}
		receiver.back = receiver.front
	} else {
		receiver.size++
		receiver.back.next = &pElement{
			previous: receiver.back,
			next:     nil,
			value:    value,
			priority: priority,
		}
		receiver.back = receiver.back.next
	}
	receiver.makeQueueOfPriotiry()
}

func (receiver *pQueue) dequeue() interface{} {
	if receiver.size == 0 {
		return nil
	} else if receiver.size == 1 {
		currentValue := receiver.front.value
		receiver.front, receiver.back = nil, nil
		receiver.size--
		return currentValue
	} else {
		currentValue := receiver.front.value
		receiver.front = receiver.front.next
		receiver.front.previous = nil
		receiver.size--
		return currentValue
	}
}

func main() {}
