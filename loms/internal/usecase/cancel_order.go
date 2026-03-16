package usecase

func (s *LOMSService) CancelOrder(orderId TOrderId) error {
	var err error

	orderInfo, err := s.orderRepo.Info(orderId)
	if err != nil {
		return err
	}

	err = s.stockRepo.ReserveCancel(&ItemCountListDTO{Items: orderInfo.Items})
	if err != nil {
		return err
	}

	err = s.orderRepo.SetState(orderId, OrderStateCancelled)
	if err != nil {
		return err
	}

	return nil

}
