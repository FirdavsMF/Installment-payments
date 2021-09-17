package billing

func Calculate(amount int, number int,  monthe int,  percent int ) int {
   
	result := ((amount / 100) * percent )
	d := result / monthe
	return d
}

func Calculate1(amount int ) int {
    percents := 6 
	result := (amount / 100) * percents + amount
	return result
}

func Calculate2(amount int, number int,  monthe int,) int {
    percent := 6
	result := ((amount / 100) * percent )+ amount
	d := result / monthe
	return d
}

func Calculate3(amount int, number int,  monthe int ) int {
    percent := 10
	result := ((amount / 100) * percent )+ amount
	d := result / monthe
	return d
}

func Calculate4(amount int, number int,  monthe int ) int {
    percent := 15
	result := ((amount / 100) * percent )+ amount
	d := result / monthe
	return d
}

func Calculate5(amount int, number int,  monthe int ) int {
    percent := 20
	result := ((amount / 100) * percent )+ amount
	d := result / monthe
	return d
}

func Calculate6(amount int, number int,  monthe int ) int {
    percent := 27
	result := ((amount / 100) * percent )+ amount
	d := result / monthe
	return d
}

func Calculate7(amount int ) int {
    percents := 10 
	result := (amount / 100) * percents + amount
	return result
}

func Calculate8(amount int ) int {
    percents := 15
	result := (amount / 100) * percents + amount
	return result
}

func Calculate9(amount int ) int {
    percents := 20 
	result := (amount / 100) * percents + amount
	return result
}

func Calculate10(amount int ) int {
    percents := 27 
	result := (amount / 100) * percents + amount
	return result
}