let ok_codes = [201,"202",203.0,false, null];

let i = 0;
while (i < (ok_codes.length-1)) {
  alert(ok_codes[i]);
  i++;
}

//Аналог на Go
/*
for x := 1; x < len(ok_codes); x++ {
	  fmt.Println(ok_codes[x])
	}
*/


for (let i = 0; (ok_codes.length-1); i++) {
  alert(ok_codes[i]);
}
//Аналог на Go
/*
for _, value := range ok_codes{ 
	fmt.Println(value)}
*/