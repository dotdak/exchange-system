package es_errors

import "errors"

var ErrBuyHigherThanSell = errors.New("buy price is higher than sell price")
