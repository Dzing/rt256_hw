package usecase

func (s *LOMSService) PayOrder(orderId TOrderId) error {
	orderInfo, err := s.orderRepo.Info(orderId)
	if err != nil {
		return err
	}

	err = s.stockRepo.ReserveRemove(&ItemCountListDTO{Items: orderInfo.Items})
	if err != nil {
		return err
	}

	err = s.orderRepo.SetState(orderId, OrderStatePayed)
	if err != nil {
		return err
	}

	return nil
}
