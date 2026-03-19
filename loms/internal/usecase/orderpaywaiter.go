package usecase

import (
	"log"
	"time"
)

type (
	eventHandler interface {
		CancelOrder(orderId TOrderId) error
	}

	OrderPayWaiter struct {
		h       eventHandler
		Timeout time.Duration
		timers  map[TOrderId]*time.Timer
	}
)

func (w *OrderPayWaiter) SetHandler(h eventHandler) {
	w.h = h
}

func (w *OrderPayWaiter) New(orderId TOrderId) *time.Timer {
	_, ok := w.timers[orderId]
	if ok {
		return nil
	}

	newTimer := time.AfterFunc(w.Timeout, func() {
		if err := w.h.CancelOrder(orderId); err != nil {
			log.Println(err)
		}
		w.Delete(orderId)
	})
	w.timers[orderId] = newTimer

	return newTimer
}

func (w *OrderPayWaiter) Get(orderId TOrderId) *time.Timer {
	return w.timers[orderId]
}

func (w *OrderPayWaiter) Stop(orderId TOrderId) {
	timer, ok := w.timers[orderId]
	if ok {
		timer.Stop()
		w.Delete(orderId)
	}
}

func (w *OrderPayWaiter) Delete(orderId TOrderId) {
	_, ok := w.timers[orderId]
	if ok {
		delete(w.timers, orderId)

	}
}

func NewOrderPayTimeWaiter() *OrderPayWaiter {
	return &OrderPayWaiter{
		Timeout: 120 * time.Second,
		timers:  make(map[TOrderId]*time.Timer),
	}
}
