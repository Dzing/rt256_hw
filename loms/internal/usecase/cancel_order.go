package usecase

import "fmt"

func (s *LOMSService) CancelOrder(orderId TOrderId) error {
	orderInfo, err := s.orderRepo.Info(orderId)
	if err != nil {
		return err
	}

	if orderInfo.OrderState != OrderStateAwaitingPayment {
		// Отменить не получится.
		return fmt.Errorf("unable to cancel order. OrderId=%v", orderId)
	}

	if err := s.stockRepo.ReserveCancel(&ItemCountListDTO{Items: orderInfo.Items}); err != nil {
		return err
	}

	if err := s.orderRepo.SetState(orderId, OrderStateCancelled); err != nil {
		return err
	}

	return nil

}
