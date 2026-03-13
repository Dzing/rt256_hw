package usecase

func (this *LOMSService) CancelOrder(orderId TOrderId) error {

	var err error

	orderInfo, err := this.orderRepo.Info(orderId)
	if err != nil {
		return err
	}

	err = this.stockRepo.ReserveCancel(&ItemCountListDTO{Items: orderInfo.Items})
	if err != nil {
		return err
	}

	err = this.orderRepo.SetState(orderId, OrderStateCancelled)
	if err != nil {
		return err
	}

	return nil

}
