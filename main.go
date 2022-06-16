package main
import "fmt"

func main() {
	var priceApple float64 = 5.99
	var pricePear  int     = 7
	var wallet     float64 = 23
	
	fmt.Println("\nЦена одного яблока - ", priceApple, "грн. Цена одной груши - ", pricePear, "грн.")
	fmt.Println("У нас в кошельке - ", wallet, " грн.", "\n")

	fmt.Println("1. Сколько денег нужно для покупки 9 яблок и 8 груш?")
	var res1 float64 = (9 * priceApple) + (8 * float64(pricePear))
	fmt.Println("Щоб купити 9 яблук та 8 груш потрібно", res1, "\n")

	fmt.Println("2. Сколько груш мы можем купить?")
	var res2 int = int(wallet) / int(pricePear)
	fmt.Println("У нас есть возможность купить ", res2, "груш.", "\n")

	fmt.Println("3. Сколько яблок мы можем купить?")
	var res3 int = int(wallet) / int(priceApple)
	fmt.Println("Мы можем купить ", res3, "яблук.", "\n")

	fmt.Println("4. Можем ли мы купить 2 яблока и 2 груши?")
	var summ = (2 * priceApple) + (2 * float64(pricePear))
	if summ < wallet {
		fmt.Println("Да")
	} else {
		fmt.Println("Нет")
	}

}
