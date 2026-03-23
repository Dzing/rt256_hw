package usecase

func (s *LOMSService) PayOrder(orderId TOrderId) error {
	orderInfo, err := s.orderRepo.Info(orderId)
	if err != nil {
		return err
	}

	if !CanChangeToOrderState(OrderStatePayed, orderInfo.OrderState) {
		// Отменить не получится.
		return &OrderStateMismatchError{OrderId: orderId, State: orderInfo.OrderState}
	}

	if err := s.stockRepo.ReserveRemove(&ItemCountListDTO{Items: orderInfo.Items}); err != nil {
		return err
	}

	if err = s.orderRepo.SetState(orderId, OrderStatePayed); err != nil {
		return err
	}

	return nil
}
