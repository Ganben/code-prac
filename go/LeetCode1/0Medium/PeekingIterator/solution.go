package PeekingIterator

type PeekingIterator struct {
	iterator      *Iterator
	hasPeeked     bool
	peekedElement interface{}
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iterator: iter, peekedElement: nil}
}

func (this *PeekingIterator) hasNext() bool {
	return this.hasPeeked || this.iterator.hasNext()
}

func (this *PeekingIterator) next() int {
	if !this.hasPeeked {
		return this.iterator.next()
	}
	result := this.peekedElement
	this.hasPeeked = false
	this.peekedElement = nil
	return result.(int)
}

func (this *PeekingIterator) peek() int {
	if !this.hasPeeked {
		this.peekedElement = this.iterator.next()
		this.hasPeeked = true
	}
	return this.peekedElement.(int)
}
