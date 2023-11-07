package composer

import (
	"context"
	"math"
	"strconv"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/pkg/slice"
)

func (c *composer) GetAllSubscriptions(ctx context.Context, userTelegramID int) (GetSubscriptionsPayload, error) {

	subs, err := c.subsMan.GetAll(ctx, userTelegramID)
	if err != nil {
		return GetSubscriptionsPayload{}, err
	}

	annual := slice.Filter(subs, func(s models.Subscription) bool {
		return s.PaymentInterval == "annual"
	})

	monthly := slice.Filter(subs, func(s models.Subscription) bool {
		return s.PaymentInterval == "monthly"
	})

	var (
		annualSum  float64
		monthlySum float64
	)

	for _, w := range annual {
		temp, err := strconv.ParseFloat(w.Amount, 64)
		if err != nil {
			return GetSubscriptionsPayload{}, err
		}

		annualSum += temp
	}

	for _, w := range monthly {
		temp, err := strconv.ParseFloat(w.Amount, 64)
		if err != nil {
			return GetSubscriptionsPayload{}, err
		}

		monthlySum += temp
	}

	annualTotal := int(math.Ceil(annualSum))
	monthlyTotal := int(math.Ceil(monthlySum))

	res := GetSubscriptionsPayload{
		MonthlyTotal: monthlyTotal,
		MonthlySubs:  monthly,
		AnnualTotal:  annualTotal,
		AnnualSubs:   annual,
	}

	return res, nil
}
