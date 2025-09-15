package context

import (
	"context"
	"fmt"
	"time"
)

func Main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	chCookPho := make(chan string)
	chCookCom := make(chan string)
	go cookCom(ctx, chCookCom)
	go cookPho(ctx, chCookPho)

	for i := 1; i <= 2; i++ {
		select {
		case pho := <-chCookPho:
			fmt.Println("Nhan duoc: ", pho)

		case com := <-chCookCom:
			fmt.Println("Nhan duoc: ", com)
		case <-ctx.Done():
			fmt.Println("Co the nhan mon")
		}
	}

}

func cookCom(ctx context.Context, ch chan<- string) {
	fmt.Println("Bat dau nau com")
	select {
	case <-time.After(2 * time.Second):
		ch <- "Com da nau xong"

	case <-ctx.Done():
		fmt.Println("Huy nau com")
		return

	}

}
func cookPho(ctx context.Context, ch chan<- string) {
	fmt.Println("Bat dau nau pho")
	select {
	case <-time.After(1 * time.Second):
		ch <- "Pho da nau xong"

	case <-ctx.Done():
		fmt.Println("Huy nau pho")
		return

	}
}

func employee(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cong viec bij huy", ctx.Err())
			return
		default:
			priority := ctx.Value("priority")
			fmt.Println("Dang lam viec: ", priority)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func exampleContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, "priority", "hight")
	go employee(ctx)

	time.Sleep(3 * time.Second)

}
