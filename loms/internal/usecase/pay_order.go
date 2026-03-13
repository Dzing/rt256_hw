package usecase

func (this *LOMSService) PayOrder(orderId TOrderId) error {

	var err error

	orderInfo, err := this.orderRepo.Info(orderId)
	if err != nil {
		return err
	}

	err = this.stockRepo.ReserveRemove(&ItemCountListDTO{Items: orderInfo.Items})
	if err != nil {
		return err
	}

	err = this.orderRepo.SetState(orderId, OrderStatePayed)
	if err != nil {
		return err
	}

	return nil

}
