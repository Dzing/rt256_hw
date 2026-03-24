package usecase

func (s *LOMSService) CancelOrder(orderId TOrderId) error {
	orderInfo, err := s.orderRepo.Info(orderId)
	if err != nil {
		return err
	}

	if !CanChangeToOrderState(OrderStateCancelled, orderInfo.OrderState) {
		// Отменить не получится.
		return &OrderStateMismatchError{OrderId: orderId, State: orderInfo.OrderState}
	}

	if err := s.stockRepo.ReserveCancel(&ItemCountListDTO{Items: orderInfo.Items}); err != nil {
		return err
	}

	if err := s.orderRepo.SetState(orderId, OrderStateCancelled); err != nil {
		return err
	}

	// Остановка таймера автоотмены заказа.
	s.payWaiter.Stop(orderId)

	return nil

}
