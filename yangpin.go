package main

import (
      "fmt"
	  "strconv"
	  "github.com/garyburd/redigo/redis"
	    )
		
	var l, m, n, q, bb int=0, 0, 0, 0, 0
	var x11, x22, x33 float64
	
func main() {
     var username, password, flag, realxs, xs, Lamdas, newBalances string
     var state, qqq int
	 var L, Lamda, newBalance float64=0, 0, 5
	 var realx, x float64
	
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
    if err != nil {
        fmt.Println("Connect to redis error", err)
        return
    }
    defer c.Close()
	
	for qqq=0;qqq<10;qqq++{
	fmt.Println("Please enter your username: ")
    fmt.Scanln(&username)
    fmt.Println("Please enter your password: ")
    fmt.Scanln(&password)
    flag, Lamda, state=identity(username, password, L)
	fmt.Println(flag)
	L = Lamda
	switch{
		case state == 1 :
	   realx, x, realxs, xs=redis1(L)
	    _, err = c.Do("SET", realxs, realx)
	    _, err = c.Do("SET", xs, x)
		x11=x
		
	   if err != nil {
        fmt.Println("redis set failed:", err)}
		
		case state == 2 :
		realx, x, realxs, xs=redis2(L)
	    _, err = c.Do("SET", realxs, realx)
	    _, err = c.Do("SET", xs, x)
		x22=x
		
	   if err != nil {
        fmt.Println("redis set failed:", err)}
		
		case state == 3 :
		realx, x, realxs, xs=redis3(L)
	    _, err = c.Do("SET", realxs, realx)
	    _, err = c.Do("SET", xs, x)
		x33=x
		
	   if err != nil {
        fmt.Println("redis set failed:", err)}
		
		case state== 4 :
		Lamda, newBalance, Lamdas, newBalances = redisL(x11, x22, x33, L)
	    _, err = c.Do("SET", Lamdas, Lamda)
	    _, err = c.Do("SET", newBalances, newBalance)
	   if err != nil {
        fmt.Println("redis set failed:", err)
	                 }
					
		 }
		if bb == 1{
			break
		}
		}
		}
	
	
     
     


func identity (u string, p string, L float64) (string, float64, int) {
	var flagx string
	var oo int
	 switch{
                              case (u =="user1" && p =="password1"):
							 oo=1
							flagx ="x1 has been saved"
							   
							 case (u =="user2" && p =="password2"):
							 oo=2
							flagx ="x2 has been saved"
							   
							 case (u =="user3" && p =="password3"):
							oo=3
							 flagx ="x3 has been saved"
							   
							 default:
							flagx = "permission denied"
							oo=0
							   
							}
	 if (l == m && m == n){
		oo=4
		L, newBalance = coordinator(x11, x22, x33, L)
		if(bb == 1){
			flagx=" Lamda has been found"
		}else{
		    flagx="Lamda & newBalance have been saved" 
	       }
	 }
	return flagx, L, oo
	
}



func abs(f float64) (float64){
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


func redis1(L float64)(float64, float64, string, string) {
	    var realx, x float64
        realx, x=optimizer1(L)
		realx1 :="realx" + strconv.Itoa(l)
		x1 := "x1" + strconv.Itoa(l)
	return realx, x, realx1, x1
	}
	
	func redis2(L float64) (float64, float64, string, string){
	    var realx, x float64
		realx, x=optimizer2(L)
		realx2 :="realx" + strconv.Itoa(m)
		x2 := "x2" + strconv.Itoa(m)
		return realx, x, realx2, x2
	}
	
	func redis3(L float64) (float64, float64, string, string){
		var realx, x float64
		realx, x=optimizer3(L)
		realx3 :="realx" + strconv.Itoa(n)
		x3 := "x3" + strconv.Itoa(n)
    	               
		return realx, x, realx3, x3
	}
	
	func redisL(x1 float64, x2 float64, x3 float64, Lamda float64) (float64, float64, string, string){
		var newBalance float64
		Lamda, newBalance = coordinator(x1, x2, x3, Lamda)
		Lamdas :="Lamda" + strconv.Itoa(q)
		newBalances := "newBalance" + strconv.Itoa(q)
    	               
		return Lamda, newBalance, Lamdas, newBalances
	}
	

func optimizer1(L float64)(float64, float64){
     var x, a float64
                          x=(4-L)*0.5
						 a=-0.5*L+2
						 l=l+1
                          return x, a
                              }

func optimizer2(L float64)(float64, float64){
     var x, a float64
                          x=(-10-L)*0.5
						 a=-L*0.5-2.5
						 m=m+1					
                          return x, a
                              }

func optimizer3(L float64)(float64, float64){
     var x, a float64
                          x=-L*0.5
						 a=-L*0.5
						 n=n+1
                          return x, a
                              }


func coordinator(x1 float64, x2 float64, x3 float64, Lamda float64)(float64, float64){
	var newBalance float64

	
						 newBalance = x1 + x2 + x3
                          Lamda = Lamda + newBalance*0.5
						 q=q+1
						 switch{
							case abs(newBalance) < 0.0000018316:
							bb = 1
							default:
							bb = 0
								}
                          return Lamda, newBalance
                              }