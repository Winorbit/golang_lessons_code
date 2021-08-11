ok_codes = [201,202,203,204,205]

# Вот этот код:
x = 0
while x < len(ok_codes):
	print(ok_codes[x])
	x += 1
	pass

"""
Аналогичен вот этому, на Go:

for x := 1; x < len(ok_codes); x++ {
	  fmt.Println(ok_codes[x])
	}
"""

# А вот этот:

for x in ok_codes:
	print(x)


"""
Вот этому коду на Go:

for _, value := range ok_codes{ 
	fmt.Println(value)}

"""