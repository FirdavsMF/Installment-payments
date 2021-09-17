package billing

func Calculate(amount float64, number float64,  monthe float64,  percent float64 ) float64 {
   
	result := ((amount / 100.0) * percent )
	d := result / monthe
	return d
}

func Calculate1(amount float64 ) float64 {
    percents := 6.0 
	result := (amount / 100.0) * percents + amount
	return result
}

func Calculate2(amount float64, number float64,  monthe float64,) float64 {
    percent := 6.0
	result := ((amount / 100.0) * percent )+ amount
	d := result / monthe
	return d
}

func Calculate3(amount float64, number float64,  monthe float64 ) float64 {
    percent := 10.0
	result := ((amount / 100.0) * percent )+ amount
	d := result / monthe
	return d
}

func Calculate4(amount float64, number float64,  monthe float64 ) float64 {
    percent := 15.0
	result := ((amount / 100.0) * percent )+ amount
	d := result / monthe
	return d
}

func Calculate5(amount float64, number float64,  monthe float64 ) float64 {
    percent := 20.0
	result := ((amount / 100.0) * percent )+ amount
	d := result / monthe
	return d
}

func Calculate6(amount float64, number float64,  monthe float64 ) float64 {
    percent := 27.0
	result := ((amount / 100.0) * percent )+ amount
	d := result / monthe
	return d
}

func Calculate7(amount float64 ) float64 {
    percents := 10.0
	result := (amount / 100.0) * percents + amount
	return result
}

func Calculate8(amount float64 ) float64 {
    percents := 15.0
	result := (amount / 100.0) * percents + amount
	return result
}

func Calculate9(amount float64 ) float64 {
    percents := 20.0
	result := (amount / 100.0) * percents + amount
	return result
}

func Calculate10(amount float64 ) float64 {
    percents := 27.0 
	result := (amount / 100.0) * percents + amount
	return result
}