package dao

import (
	"math"

	"github.com/go-playground/validator/v10"
)

var validate = func() *validator.Validate {
	validate := validator.New()
	validate.RegisterStructValidation(WagerStructLevelValidation, Wager{})
	validate.RegisterStructValidation(WagerStructLevelValidation, Wager{})
	return validate
}()

func WagerStructLevelValidation(sl validator.StructLevel) {
	wager := sl.Current().Interface().(Wager)

	if wager.SellingPrice*1e2-math.Floor(wager.SellingPrice*1e2) > 1e-8 {
		sl.ReportError(wager.SellingPrice, "sprice", "SellingPrice", "spricemonetaryformat", "")
		return
	}

	if wager.SellingPrice <= float64(wager.TotalWagerValue*wager.SellingPercentage)/100 {
		// sl.ReportError("field SellingPrice must be larger than TotalWagerValue * SellingPercentage")
		sl.ReportError(wager.SellingPrice, "sprice", "SellingPrice", "spricegttotal", "")
		sl.ReportError(wager.TotalWagerValue, "tvalue", "TotalValue", "spricegttotal", "")
		sl.ReportError(wager.SellingPercentage, "spercent", "SellingPercentage", "spricegttotal", "")
		return
	}

	if (wager.AmountSold == nil || *wager.AmountSold <= 0) && wager.CurrentSellingPrice > wager.SellingPrice {
		sl.ReportError(wager.CurrentSellingPrice, "csprice", "CurrentSellingPrice", "currentgtsellingprice", "")
		return
	}

	if wager.AmountSold != nil && *wager.AmountSold > 0 && *wager.AmountSold+wager.CurrentSellingPrice > wager.SellingPrice {
		sl.ReportError(wager.CurrentSellingPrice, "csprice", "CurrentSellingPrice", "amountcurrentgtsellingprice", "")
		return
	}
}
